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

type EmploymentHandler struct {
	repo *repositories.EmploymentRepository
}

func NewEmploymentHandler(repo *repositories.EmploymentRepository) *EmploymentHandler {
	return &EmploymentHandler{repo: repo}
}

func (h *EmploymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Get user_id from the URL
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
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

	var employment models.Employment
	if err := json.NewDecoder(r.Body).Decode(&employment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign user_id from the URL to the employment record
	employment.UserID = uint(userID)

	// Parse dates
	if employment.RawStartDate != "" {
		startDate, err := time.Parse("02-01-2006", employment.RawStartDate)
		if err != nil {
			http.Error(w, "Invalid startDate format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		employment.StartDate = startDate
	}
	if employment.RawEndDate != "" {
		endDate, err := time.Parse("02-01-2006", employment.RawEndDate)
		if err != nil {
			http.Error(w, "Invalid endDate format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		employment.EndDate = endDate
	}

	// Save the employment record
	if err := h.repo.Create(&employment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employment)
}

func (h *EmploymentHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	// Get user_id from the URL
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
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

	// Fetch employment records for the user
	employments, err := h.repo.GetByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch employment records", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employments)
}

func (h *EmploymentHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from the URL path
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Extract id from the query parameter
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}

	employmentID, err := strconv.Atoi(ids[0])
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

	// Decode request body
	var employment models.Employment
	if err := json.NewDecoder(r.Body).Decode(&employment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate if employment exists for the given user_id
	existingEmployment, err := h.repo.GetByID(uint(employmentID))
	if err != nil || existingEmployment.UserID != uint(userID) {
		http.Error(w, "Employment not found for the specified user", http.StatusNotFound)
		return
	}

	// Update employment record
	employment.ID = uint(employmentID)
	employment.UserID = uint(userID)

	// Parse dates
	if employment.RawStartDate != "" {
		startDate, err := time.Parse("02-01-2006", employment.RawStartDate)
		if err != nil {
			http.Error(w, "Invalid startDate format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		employment.StartDate = startDate
	}
	if employment.RawEndDate != "" {
		endDate, err := time.Parse("02-01-2006", employment.RawEndDate)
		if err != nil {
			http.Error(w, "Invalid endDate format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		employment.EndDate = endDate
	}

	if err := h.repo.Update(&employment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employment)
}

func (h *EmploymentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from the URL path
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Extract id from the query parameter
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}

	employmentID, err := strconv.Atoi(ids[0])
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

	// Validate if employment exists for the given user_id
	employment, err := h.repo.GetByID(uint(employmentID))
	if err != nil || employment.UserID != uint(userID) {
		http.Error(w, "Employment not found for the specified user", http.StatusNotFound)
		return
	}

	// Delete the employment record
	if err := h.repo.Delete(uint(employmentID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
