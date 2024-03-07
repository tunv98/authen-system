package models

import (
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	gorm.Model
	CampaignName  string `gorm:"unique"`
	TotalVouchers uint   `gorm:"not null"`
	DiscountValue float64
	StartDate     time.Time
	EndDate       time.Time
}
