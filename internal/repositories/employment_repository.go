package repositories

import (
	"github.com/raviqlahadi/cv-form-backend/internal/db"
	"github.com/raviqlahadi/cv-form-backend/internal/models"
	"gorm.io/gorm"
)

type EmploymentRepository struct {
	db *gorm.DB
}

func NewEmploymentRepository() *EmploymentRepository {
	return &EmploymentRepository{db: db.DB}
}

func (r *EmploymentRepository) Create(employment *models.Employment) error {
	return r.db.Create(employment).Error
}

func (r *EmploymentRepository) GetByID(id uint) (*models.Employment, error) {
	var employment models.Employment
	if err := r.db.First(&employment, id).Error; err != nil {
		return nil, err
	}
	return &employment, nil
}

func (r *EmploymentRepository) GetByUserID(userID uint) ([]models.Employment, error) {
	var employments []models.Employment
	err := r.db.Where("user_id = ?", userID).Find(&employments).Error
	if err != nil {
		return nil, err
	}
	return employments, nil
}

func (r *EmploymentRepository) Update(employment *models.Employment) error {
	return r.db.Save(employment).Error
}

func (r *EmploymentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Employment{}, id).Error
}

func (r *EmploymentRepository) CheckUserExists(userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *EmploymentRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Employment{}).Error
}
