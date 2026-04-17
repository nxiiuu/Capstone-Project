package entity

import (
	"gorm.io/gorm"
)


type ProjectRepositoryStatus struct {
	gorm.Model
	StatusName string `gorm:"not null" json:"status_name"`

	ProjectRepositories []ProjectRepository `gorm:"foreignKey:StatusID" json:"project_repositories"`
}
