package entity

import (
	"gorm.io/gorm"
)


type FinalScore struct {
	gorm.Model
	TotalScore float64 `gorm:"type:decimal(5,2);not null" json:"total_score"`
	Grade      string  `gorm:"not null" json:"grade"`

	UserID   *uint `gorm:"not null" json:"user_id"`
	GradeID  *uint `json:"grade_id"`
	SGroupID *uint `gorm:"not null" json:"s_group_id"`

	GradeItem    *Grade        `gorm:"foreignKey:GradeID" json:"grade_item"`
	StudentGroup *StudentGroup `gorm:"foreignKey:SGroupID" json:"student_group"`
}
