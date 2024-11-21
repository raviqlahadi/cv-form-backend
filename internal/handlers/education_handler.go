package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/raviqlahadi/cv-form-backend/internal/models"
	"github.com/raviqlahadi/cv-form-backend/internal/repositories"
)

type EducationHandler struct {
	repo *repositories.EducationRepository
}

func NewEducationHandler(repo *repositories.EducationRepository) *EducationHandler {
	return &EducationHandler{repo: repo}
}

func (h *EducationHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	var education models.Education
	if err := json.NewDecoder(r.Body).Decode(&education); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	education.UserID = uint(userID)

	// Parse dates
	if education.RawStartDate != "" {
		startDate, err := time.Parse("02-01-2006", education.RawStartDate)
		if err != nil {
			http.Error(w, "Invalid startDate format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		education.StartDate = startDate
	}
	if education.RawEndDate != "" {
		endDate, err := time.Parse("02-01-2006", education.RawEndDate)
		if err != nil {
			http.Error(w, "Invalid endDate format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		education.EndDate = endDate
	}

	if err := h.repo.Create(&education); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(education)
}

func (h *EducationHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	educations, err := h.repo.GetByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch educations", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(educations)
}

func (h *EducationHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}

	educationID, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid id query parameter", http.StatusBadRequest)
		return
	}

	var education models.Education
	if err := json.NewDecoder(r.Body).Decode(&education); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	education.ID = uint(educationID)
	education.UserID = uint(userID)

	if education.RawStartDate != "" {
		startDate, err := time.Parse("02-01-2006", education.RawStartDate)
		if err != nil {
			http.Error(w, "Invalid startDate format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		education.StartDate = startDate
	}
	if education.RawEndDate != "" {
		endDate, err := time.Parse("02-01-2006", education.RawEndDate)
		if err != nil {
			http.Error(w, "Invalid endDate format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		education.EndDate = endDate
	}

	if err := h.repo.Update(&education); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(education)
}

func (h *EducationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}

	educationID, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid id query parameter", http.StatusBadRequest)
		return
	}

	// Validate if user_id exists in the users table
	exists, err := h.repo.CheckUserExists(uint(userID))
	if err != nil {
		http.Error(w, "Database error while validating user_id", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := h.repo.Delete(uint(educationID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
