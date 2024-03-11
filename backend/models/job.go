package models

import (
	"time"

	"gorm.io/gorm"
)

const (
    EmploymentFullTime = "Full time"
    EmploymentPartTime = "Part time"
    EmploymentRemote   = "Remote"
)

type Job struct {
    gorm.Model
    Title            string
    Description      string
    ExperienceLevel  string `gorm:"column:experience_level"`
    DatePosted       time.Time `gorm:"column:date_posted"`
    DateExpiry       time.Time `gorm:"column:date_expiry"`
    Vacancies        int
    CompanyID        uint `gorm:"column:company_id"`
    Location         string
    EmploymentType   string `gorm:"column:employment_type;type:varchar(255);default:'Full time'"`
    Applications     []Application `gorm:"foreignKey:JobID"`
    Bookmarks        []Bookmark    `gorm:"foreignKey:JobID"`
    Sectors          []Sector      `gorm:"many2many:job_sectors;"`
}