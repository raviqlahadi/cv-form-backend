package repositories

import (
	"github.com/raviqlahadi/cv-form-backend/internal/db"
	"github.com/raviqlahadi/cv-form-backend/internal/models"
	"gorm.io/gorm"
)

type SkillRepository struct {
	db *gorm.DB
}

func NewSkillRepository() *SkillRepository {
	return &SkillRepository{db.DB}
}

func (r *SkillRepository) Create(skill *models.Skill) error {
	return r.db.Create(skill).Error
}

func (r *SkillRepository) GetByID(id uint) (*models.Skill, error) {
	var skill models.Skill
	err := r.db.First(&skill, id).Error
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func (r *SkillRepository) GetByUserID(userID uint) ([]models.Skill, error) {
	var skills []models.Skill
	err := r.db.Where("user_id = ?", userID).Find(&skills).Error
	return skills, err
}

func (r *SkillRepository) Update(skill *models.Skill) error {
	return r.db.Save(skill).Error
}

func (r *SkillRepository) Delete(id uint) error {
	return r.db.Delete(&models.Skill{}, id).Error
}

func (r *SkillRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Skill{}).Error
}

func (r *SkillRepository) CheckUserExists(userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
