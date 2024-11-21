package repositories

import (
	"github.com/raviqlahadi/cv-form-backend/internal/db"
	"github.com/raviqlahadi/cv-form-backend/internal/models"
	"gorm.io/gorm"
)

type EducationRepository struct {
	db *gorm.DB
}

func NewEducationRepository() *EducationRepository {
	return &EducationRepository{db.DB}
}

func (r *EducationRepository) Create(education *models.Education) error {
	return r.db.Create(education).Error
}

func (r *EducationRepository) GetByID(id uint) (*models.Education, error) {
	var education models.Education
	err := r.db.First(&education, id).Error
	if err != nil {
		return nil, err
	}
	return &education, nil
}

func (r *EducationRepository) GetByUserID(userID uint) ([]models.Education, error) {
	var educations []models.Education
	err := r.db.Where("user_id = ?", userID).Find(&educations).Error
	return educations, err
}

func (r *EducationRepository) Update(education *models.Education) error {
	return r.db.Save(education).Error
}

func (r *EducationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Education{}, id).Error
}

func (r *EducationRepository) CheckUserExists(userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *EducationRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Education{}).Error
}
