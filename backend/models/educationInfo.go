package models

import (
	"time"

	"gorm.io/gorm"
)

type EducationInfo struct {
    gorm.Model
    LearningTitle    string `gorm:"column:learning_title"`
    SchoolUniversity string `gorm:"column:school_university"`
    StartDate        time.Time `gorm:"column:start_date"`
    EndDate          time.Time `gorm:"column:end_date"`
    Description      string
    UserID           uint `gorm:"column:user_id"`
}