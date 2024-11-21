package models

import (
	"time"
)

type Skill struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"column:user_id" json:"userId"`
	Skill     string    `gorm:"column:skill" json:"skill"`
	Level     string    `gorm:"column:level" json:"level"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
