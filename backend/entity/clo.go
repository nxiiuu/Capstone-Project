package entity

import (
	"gorm.io/gorm"
)


type CLO struct {
	gorm.Model
	CLODescription string `json:"clo_description"`
	Weight         int    `json:"weight"`

	Rubrics []Rubric `gorm:"foreignKey:CLOID" json:"rubrics"`
}
