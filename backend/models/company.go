package models

import (
	"gorm.io/gorm"
)

type Company struct {
    gorm.Model
    Name              string `gorm:"column:name" binding:"required" json:"name"`
    RegistrationEmail string `gorm:"unique" binding:"required" json:"email_address"`
    Password          string `gorm:"column:password" binding:"required" json:"password"`
    Location          string
    Sector            string
    NoOfEmployees     int
    Logo              string
    City              string
    Region            string
    SubCityZone       string `gorm:"column:sub_city_zone"`
    WebsiteLink       string
    PhoneNumber       string `gorm:"column:phone_number" binding:"required" json:"phone_number"`
    EmailAddress      string `gorm:"unique"`
    Jobs              []Job `gorm:"foreignKey:CompanyID"`
}