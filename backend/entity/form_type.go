package entity

import (
	"gorm.io/gorm"
)

type FormType struct {
	gorm.Model
	TypeName string `json:"type_name"`

	AssessmentForms []AssessmentForm `gorm:"foreignKey:FormTypeID" json:"assessment_forms"`
}
