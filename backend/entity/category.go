package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `gorm:"not null" json:"category_name"`

	ProjectRepositories []ProjectRepository `gorm:"foreignKey:CategoryID" json:"project_repositories"`
}
