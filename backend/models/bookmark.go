package models

import (
	"gorm.io/gorm"
)

type Bookmark struct {
    gorm.Model
    JobID  uint `gorm:"column:job_id"`
    UserID uint `gorm:"column:user_id"`
    UniqueConstraint string `gorm:"uniqueIndex:idx_job_user"`
}