package models

import (
	"time"

	"gorm.io/gorm"
)

type JobSeeker struct {
    gorm.Model
    Name              string `gorm:"column:name" binding:"required" json:"name"`
    Password          string `gorm:"column:password" binding:"required" json:"password"`
    Gender            string `gorm:"column:gender"  json:"gender"`
    DoB               time.Time `gorm:"column:dob"  json:"dob"`
    City              string `gorm:"column:city" json:"city"`
    Region            string `gorm:"column:region"  json:"region"`
    SubCityZone       string `gorm:"column:sub_city_zone"  json:"sub_city_zone"`
    PortfolioWebsite  string `gorm:"column:portfolio_website"`
    LinkedIn          string `gorm:"column:linked_in"  json:"linked_in"`
    PhoneNumber       string `gorm:"column:phone_number" binding:"required" json:"phone_number"`
    EmailAddress      string `gorm:"unique" binding:"required" json:"email_address"`
    ProfessionalTitle string `gorm:"column:professional_title"`
    Bio               string `gorm:"column:bio"  json:"bio"`
    Applications      []Application `gorm:"foreignKey:UserID"`
    Bookmarks         []Bookmark    `gorm:"foreignKey:UserID"`
    EducationInfos    []EducationInfo `gorm:"foreignKey:UserID"`
    WorkHistories     []WorkHistory `gorm:"foreignKey:UserID"`
}
