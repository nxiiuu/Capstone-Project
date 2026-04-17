package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	SutID     string `json:"sut_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	RoleID    *uint  `json:"role_id"`

	Role                *Role               `gorm:"foreignKey:RoleID" json:"role"`
	ProjectRepositories []ProjectRepository `gorm:"foreignKey:UserID" json:"project_repositories"`
	Messages            []Messages          `gorm:"foreignKey:UserID" json:"messages"`
}
