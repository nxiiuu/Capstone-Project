package entity

import (
	"gorm.io/gorm"
)


type StudentGroup struct {
	gorm.Model
	ProjectTitleTH  string `json:"project_title_th"`
	ProjectTitleEN  string `json:"project_title_en"`
	ProjectAbstract string `json:"project_abstract"`
	GroupStatus     string `json:"group_status"`
	FormID          *uint  `json:"form_id"`
	UserID          *uint  `json:"user_id"`

	Results             []Result            `gorm:"foreignKey:SGroupID" json:"results"`
	ProjectRepositories []ProjectRepository `gorm:"foreignKey:SGroupID" json:"project_repositories"`
	FinalScores         []FinalScore        `gorm:"foreignKey:SGroupID" json:"final_scores"`
	Chatrooms           []Chatrooms         `gorm:"foreignKey:SGroupID" json:"chatrooms"`
}
