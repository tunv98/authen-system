package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FullName    string `gorm:"not null"`
	PhoneNumber string `gorm:"unique"`
	Email       string `gorm:"unique"`
	UserName    string `gorm:"unique"`
	PassWord    string `gorm:"not null"`
	Birthday    string
	LatestLogin time.Time `gorm:"default:null"`
}
