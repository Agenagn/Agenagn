package models

import (
	"gorm.io/gorm"
)

type Sector struct {
    gorm.Model
    Title string
}