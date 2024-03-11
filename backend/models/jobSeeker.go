package models

import (
	"time"

	"gorm.io/gorm"
)

type JobSeeker struct {
    gorm.Model
    Name              string
    Password          string
    Gender            string
    DoB               time.Time `gorm:"column:dob"`
    City              string
    Region            string
    SubCityZone       string `gorm:"column:sub_city_zone"`
    PortfolioWebsite  string `gorm:"column:portfolio_website"`
    LinkedIn          string
    GitHub            string
    PhoneNumber       string
    EmailAddress      string `gorm:"unique"`
    ProfessionalTitle string `gorm:"column:professional_title"`
    Bio               string
    Applications      []Application `gorm:"foreignKey:UserID"`
    Bookmarks         []Bookmark    `gorm:"foreignKey:UserID"`
    EducationInfos    []EducationInfo `gorm:"foreignKey:UserID"`
    WorkHistories     []WorkHistory `gorm:"foreignKey:UserID"`
}