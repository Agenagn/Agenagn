package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkHistory struct {
    gorm.Model
    Position    string
    Company     string
    StartDate   time.Time `gorm:"column:start_date"`
    EndDate     time.Time `gorm:"column:end_date"`
    Description string
    UserID      uint `gorm:"column:user_id"`
}