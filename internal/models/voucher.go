package models

import (
	"gorm.io/gorm"
	"time"
)

type VoucherStatus string

const (
	ActiveStatus  VoucherStatus = "active"
	UsedStatus    VoucherStatus = "used"
	ExpiredStatus VoucherStatus = "expired"
)

type Voucher struct {
	gorm.Model
	Code        string `gorm:"unique"`
	CampaignID  uint   `gorm:"foreignKey:CampaignID"`
	UserID      uint   `gorm:"foreignKey:UserID"`
	ExpiredTime time.Time
	Status      VoucherStatus `gorm:"type:ENUM('active','used','expired')"`
}
