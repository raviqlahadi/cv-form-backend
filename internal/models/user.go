package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	FirstName         string    `gorm:"column:first_name" json:"firstName"`
	LastName          string    `gorm:"column:last_name" json:"lastName"`
	Email             string    `gorm:"unique" json:"email"`
	Phone             string    `gorm:"column:phone" json:"phone"`
	Country           string    `gorm:"column:country" json:"country"`
	City              string    `gorm:"column:city" json:"city"`
	Address           string    `gorm:"column:address" json:"address"`
	PostalCode        int       `gorm:"column:postal_code" json:"postalCode"`
	WantedJobTitle    string    `gorm:"column:wanted_job_title" json:"wantedJobTitle"`
	DrivingLicense    string    `gorm:"column:driving_license" json:"drivingLicense"`
	Nationality       string    `gorm:"column:nationality" json:"nationality"`
	PlaceOfBirth      string    `gorm:"column:place_of_birth" json:"placeOfBirth"`
	DateOfBirth       time.Time `gorm:"column:date_of_birth" json:"-"`
	RawDateOfBirth    string    `json:"dateOfBirth" gorm:"-"`
	PhotoURL          string    `gorm:"column:photo_url" json:"photoUrl"`
	WorkingExperience string    `gorm:"column:working_experience" json:"workingExperience"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		*Alias
		DateOfBirth string `json:"dateOfBirth"`
	}{
		Alias:       (*Alias)(u),
		DateOfBirth: u.DateOfBirth.Format("02-01-2006"),
	})
}
