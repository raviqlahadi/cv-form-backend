package models

import (
	"encoding/json"
	"time"
)

type Employment struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"column:user_id" json:"userId"`
	JobTitle     string    `gorm:"column:job_title" json:"jobTitle"`
	Employer     string    `gorm:"column:employer" json:"employer"`
	StartDate    time.Time `gorm:"column:start_date" json:"-"`
	EndDate      time.Time `gorm:"column:end_date" json:"-"`
	City         string    `gorm:"column:city" json:"city"`
	Description  string    `gorm:"column:description" json:"description"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updatedAt"`
	RawStartDate string    `json:"startDate" gorm:"-"`
	RawEndDate   string    `json:"endDate" gorm:"-"`
}

// MarshalJSON formats the response for Employment
func (e *Employment) MarshalJSON() ([]byte, error) {
	type Alias Employment
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
