package repositories

import (
	"github.com/raviqlahadi/cv-form-backend/internal/db"
	"github.com/raviqlahadi/cv-form-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: db.DB}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) UpdateWorkingExperience(userID uint, workingExperience string) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("working_experience", workingExperience).Error
}

func (r *UserRepository) GetWorkingExperience(userID uint) (string, error) {
	var user models.User
	err := r.db.Select("working_experience").
		Where("id = ?", userID).
		First(&user).Error
	if err != nil {
		return "", err
	}
	return user.WorkingExperience, nil
}

func (r *UserRepository) CheckUserExists(userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) UpdatePhotoURL(userID uint, photoURL string) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("photo_url", photoURL).Error
}

func (r *UserRepository) ClearPhotoURL(userID uint) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("photo_url", "").Error
}

func (r *UserRepository) GetPhotoURL(userID uint) (string, error) {
	var user models.User
	err := r.db.Select("photo_url").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.PhotoURL, nil
}
