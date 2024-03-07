package database

import (
	"authen-system/internal/models"
	"gorm.io/gorm"
	"time"
)

type VoucherRepository interface {
	Create(req VoucherRequest) error
}

type voucherRepo struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepo{db: db}
}

type VoucherRequest struct {
	CampaignID uint
	UserID     uint
	EndDate    time.Time
}

func (r *voucherRepo) Create(req VoucherRequest) error {
	status := models.ActiveStatus
	if time.Now().Before(req.EndDate) {
		status = models.ExpiredStatus
	}
	voucher := models.Voucher{
		Code:        generateUniqueCode(),
		CampaignID:  req.CampaignID,
		UserID:      req.UserID,
		ExpiredTime: req.EndDate,
		Status:      status,
	}
	return r.db.Create(voucher).Error
}
