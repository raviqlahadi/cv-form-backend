package models

import (
	"encoding/json"
	"time"
)

type Education struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"column:user_id" json:"userId"`
	School       string    `gorm:"column:school" json:"school"`
	Degree       string    `gorm:"column:degree" json:"degree"`
	StartDate    time.Time `gorm:"column:start_date" json:"-"`
	EndDate      time.Time `gorm:"column:end_date" json:"-"`
	City         string    `gorm:"column:city" json:"city"`
	Description  string    `gorm:"column:description" json:"description"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`
	RawStartDate string    `json:"startDate" gorm:"-"`
	RawEndDate   string    `json:"endDate" gorm:"-"`
}

// MarshalJSON formats the response for Education
func (e *Education) MarshalJSON() ([]byte, error) {
	type Alias Education
	return json.Marshal(&struct {
		*Alias
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}{
		Alias:     (*Alias)(e),
		StartDate: e.StartDate.Format("02-01-2006"),
		EndDate:   e.EndDate.Format("02-01-2006"),
	})
}
