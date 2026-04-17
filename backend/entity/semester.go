package entity

import (
	"time"
	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model
	AcademicYear int       `gorm:"not null" json:"academic_year"`
	Term         int       `gorm:"not null" json:"term"`
	StartDate    time.Time `gorm:"not null" json:"start_date"`
	EndDate      time.Time `gorm:"not null" json:"end_date"`

	ProjectRepositories []ProjectRepository `gorm:"foreignKey:SemesterID" json:"project_repositories"`
	SystemSetting       *SystemSetting      `gorm:"foreignKey:SemesterID" json:"system_setting"`
	AssessmentForms     []AssessmentForm    `gorm:"foreignKey:SemesterID" json:"assessment_forms"`
}
