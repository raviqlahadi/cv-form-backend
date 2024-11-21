package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raviqlahadi/cv-form-backend/internal/models"
	"github.com/raviqlahadi/cv-form-backend/internal/repositories"
)

type SkillHandler struct {
	repo *repositories.SkillRepository
}

func NewSkillHandler(repo *repositories.SkillRepository) *SkillHandler {
	return &SkillHandler{repo: repo}
}

func (h *SkillHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Get user_id from the URL
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Check if user exists
	exists, err := h.repo.CheckUserExists(uint(userID))
	if err != nil {
		http.Error(w, "Database error while validating user_id", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var skill models.Skill
	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	skill.UserID = uint(userID)

	if err := h.repo.Create(&skill); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(skill)
}

func (h *SkillHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	exists, err := h.repo.CheckUserExists(uint(userID))
	if err != nil {
		http.Error(w, "Database error while validating user_id", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	skills, err := h.repo.GetByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch skills", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(skills)
}

func (h *SkillHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	skillID, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid id query parameter", http.StatusBadRequest)
		return
	}

	var skill models.Skill
	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	skill.ID = uint(skillID)
	skill.UserID = uint(userID)

	if err := h.repo.Update(&skill); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(skill)
}

func (h *SkillHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	skillID, err := strconv.Atoi(ids[0])
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

	// Validate if skill exists for the given user_id
	skill, err := h.repo.GetByID(uint(skillID))
	if err != nil || skill.UserID != uint(userID) {
		http.Error(w, "skill not found for the specified user", http.StatusNotFound)
		return
	}

	if err := h.repo.Delete(uint(skillID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
