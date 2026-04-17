package entity

import (
	"gorm.io/gorm"
)


type Grade struct {
	gorm.Model
	GradeName    string  `json:"grade_name"`
	MinimumScore float64 `gorm:"type:decimal(5,2)" json:"minimum_score"`

	FinalScores []FinalScore `gorm:"foreignKey:GradeID" json:"final_scores"`
}
