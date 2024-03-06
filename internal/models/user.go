package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FullName    string    `json:"fullName" gorm:"not null"`
	PhoneNumber string    `json:"phoneNumber" gorm:"uniqueIndex"`
	Email       string    `json:"email" gorm:"uniqueIndex"`
	UserName    string    `json:"userName" gorm:"uniqueIndex"`
	PassWord    string    `json:"passWord" gorm:"not null"`
	Birthday    time.Time `json:"birthday"`
	LatestLogin time.Time `json:"latestLogin"`
}
