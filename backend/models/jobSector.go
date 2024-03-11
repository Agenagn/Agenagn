package models

import (
	"gorm.io/gorm"
)

type JobSector struct {
    gorm.Model
    SectorID uint `gorm:"column:sector_id"`
    JobID    uint `gorm:"column:job_id"`
    UniqueConstraint string `gorm:"uniqueIndex:idx_sector_job"`
}