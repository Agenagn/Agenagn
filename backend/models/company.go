package models

import (
	"gorm.io/gorm"
)

type Company struct {
    gorm.Model
    Name              string
    RegistrationEmail string `gorm:"unique"`
    Password          string
    Location          string
    Sector            string
    NoOfEmployees     int
    Logo              string
    City              string
    Region            string
    SubCityZone       string `gorm:"column:sub_city_zone"`
    WebsiteLink       string
    PhoneNumber       string
    EmailAddress      string `gorm:"unique"`
    Jobs              []Job `gorm:"foreignKey:CompanyID"`
}