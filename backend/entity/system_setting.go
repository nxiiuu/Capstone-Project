package entity

import (
	"time"
	"gorm.io/gorm"
)

type SystemSetting struct {
	gorm.Model
	RegistrationOpen bool      `json:"registration_open"`
	LastUpdated      time.Time `json:"last_updated"`
	SemesterID       *uint     `json:"semester_id"`

	Semester *Semester `gorm:"foreignKey:SemesterID" json:"semester"`
}
