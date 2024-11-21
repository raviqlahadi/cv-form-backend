package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/raviqlahadi/cv-form-backend/internal/models"
	"github.com/raviqlahadi/cv-form-backend/internal/repositories"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse dateOfBirth
	if user.RawDateOfBirth != "" {
		parsedDate, err := time.Parse("02-01-2006", user.RawDateOfBirth)
		if err != nil {
			http.Error(w, "Invalid date format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		user.DateOfBirth = parsedDate
	}

	if err := h.repo.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetByID(uint(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse dateOfBirth
	if user.RawDateOfBirth != "" {
		parsedDate, err := time.Parse("02-01-2006", user.RawDateOfBirth)
		if err != nil {
			http.Error(w, "Invalid date format. Use DD-MM-YYYY.", http.StatusBadRequest)
			return
		}
		user.DateOfBirth = parsedDate
	}

	user.ID = uint(id)
	if err := h.repo.Update(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateWorkingExperience(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from the URL
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Decode the request body
	var requestBody struct {
		WorkingExperience string `json:"workingExperience"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate if user exists
	exists, err := h.repo.CheckUserExists(uint(userID))
	if err != nil {
		http.Error(w, "Database error while validating user_id", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update the working_experience
	if err := h.repo.UpdateWorkingExperience(uint(userID), requestBody.WorkingExperience); err != nil {
		http.Error(w, "Failed to update working experience", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) GetWorkingExperience(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from the URL
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Validate if user exists
	exists, err := h.repo.CheckUserExists(uint(userID))
	if err != nil {
		http.Error(w, "Database error while validating user_id", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Fetch the working_experience
	workingExperience, err := h.repo.GetWorkingExperience(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch working experience", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"workingExperience": workingExperience,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) UploadPhoto(w http.ResponseWriter, r *http.Request) {
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

	var requestBody struct {
		Base64Img string `json:"base64img"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decode Base64 string
	base64Data := requestBody.Base64Img[strings.IndexByte(requestBody.Base64Img, ',')+1:]
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		http.Error(w, "Invalid Base64 data", http.StatusBadRequest)
		return
	}

	// Save image to the ./image-upload directory
	filePath := fmt.Sprintf("./image-upload/user_%d.png", userID)
	err = os.WriteFile(filePath, imageData, 0644)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}

	photoURL := fmt.Sprintf("/image-upload/user_%d.png", userID)
	if err := h.repo.UpdatePhotoURL(uint(userID), photoURL); err != nil {
		http.Error(w, "Failed to update PhotoURL", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"photoUrl": photoURL}
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) DownloadPhoto(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	photoURL, err := h.repo.GetPhotoURL(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch photo URL", http.StatusInternalServerError)
		return
	}
	if photoURL == "" {
		http.Error(w, "No photo found", http.StatusNotFound)
		return
	}

	filePath := "." + photoURL
	http.ServeFile(w, r, filePath)
}

func (h *UserHandler) DeletePhoto(w http.ResponseWriter, r *http.Request) {
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

	photoURL, err := h.repo.GetPhotoURL(uint(userID))
	if err != nil {
		http.Error(w, "Failed to fetch photo URL", http.StatusInternalServerError)
		return
	}

	if photoURL != "" {
		filePath := "." + photoURL
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			http.Error(w, "Failed to delete image file", http.StatusInternalServerError)
			return
		}
	}

	if err := h.repo.ClearPhotoURL(uint(userID)); err != nil {
		http.Error(w, "Failed to clear PhotoURL", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
