package entity

import (
	"gorm.io/gorm"
)


type Rubric struct {
	gorm.Model
	Detail   string `json:"detail"`
	MaxScore int    `json:"max_score"`
	CLOID    *uint  `json:"clo_id"`
	FormID   *uint  `json:"form_id"`

	CLO            *CLO            `gorm:"foreignKey:CLOID" json:"clo"`
	AssessmentForm *AssessmentForm `gorm:"foreignKey:FormID" json:"assessment_form"`
	Results        []Result        `gorm:"foreignKey:RubricID" json:"results"`
}
