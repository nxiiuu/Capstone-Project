package entity

import (
	"gorm.io/gorm"
)


type Result struct {
	gorm.Model
	Score    string `json:"score"`
	Comment  string `json:"comment"`
	RubricID *uint  `json:"rubric_id"`
	SGroupID *uint  `json:"s_group_id"`

	Rubric       *Rubric       `gorm:"foreignKey:RubricID" json:"rubric"`
	StudentGroup *StudentGroup `gorm:"foreignKey:SGroupID" json:"student_group"`
}
