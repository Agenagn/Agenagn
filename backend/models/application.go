package models

import (
	"gorm.io/gorm"
)

const (
    StatusAccepted   = "Accepted"
    StatusInProgress = "In Progress"
    StatusRejected   = "Rejected"
)

type Application struct {
    gorm.Model
    JobID         uint `gorm:"column:job_id"`
    UserID        uint `gorm:"column:user_id"`
    Resume        string
    CoverLetter   string `gorm:"column:cover_letter"`
    Status        string `gorm:"column:status;type:varchar(255);default:'In Progress'"`
}