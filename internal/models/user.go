package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FullName    string    `json:"fullName" gorm:"not null"`
	PhoneNumber string    `json:"phoneNumber" gorm:"unique"`
	Email       string    `json:"email" gorm:"unique"`
	UserName    string    `json:"userName" gorm:"unique"`
	PassWord    string    `json:"passWord" gorm:"not null"`
	Birthday    string    `json:"birthday"`
	LatestLogin time.Time `json:"latestLogin"`
}
