package entity

import (
	"gorm.io/gorm"
)

type AssessmentForm struct {
	gorm.Model
	FormName   string `json:"form_name"`
	Week       string `json:"week"`
	Weight     int    `json:"weight"`
	FormTypeID *uint  `json:"form_type_id"`
	SemesterID *uint  `json:"semester_id"`

	FormType *FormType `gorm:"foreignKey:FormTypeID" json:"form_type"`
	Semester *Semester `gorm:"foreignKey:SemesterID" json:"semester"`
	Rubrics  []Rubric  `gorm:"foreignKey:FormID" json:"rubrics"`
}
