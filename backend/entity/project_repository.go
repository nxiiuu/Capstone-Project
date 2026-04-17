package entity

import (
	"gorm.io/gorm"
)

type ProjectRepository struct {
	gorm.Model
	ProjectID        string `gorm:"not null" json:"project_id"`
	ReportFile       string `gorm:"not null" json:"report_file"` // mapped from Report_File
	PresentationFile string `gorm:"not null" json:"presentation_file"` // mapped from Presentation_File
	SourceCodeLink   string `gorm:"not null" json:"source_code_link"` // mapped from Source_Code_Link

	UserID     *uint `gorm:"not null" json:"user_id"`
	CategoryID *uint `gorm:"not null" json:"category_id"`
	StatusID   *uint `gorm:"not null" json:"status_id"`
	SGroupID   *uint `gorm:"not null" json:"s_group_id"`
	SemesterID *uint `gorm:"not null" json:"semester_id"`

	User                    *User                    `gorm:"foreignKey:UserID" json:"user"`
	Category                *Category                `gorm:"foreignKey:CategoryID" json:"category"`
	ProjectRepositoryStatus *ProjectRepositoryStatus `gorm:"foreignKey:StatusID" json:"project_repository_status"`
	Semester                *Semester                `gorm:"foreignKey:SemesterID" json:"semester"`
	StudentGroup            *StudentGroup            `gorm:"foreignKey:SGroupID" json:"student_group"`
}
