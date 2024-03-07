package database

import (
	"authen-system/internal/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CampaignRepository interface {
	Create(campaign *models.Campaign) error
	GetCampaignByName(campaignName string) (models.Campaign, error)
}

type campaignRepo struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &campaignRepo{db: db}
}

func (r *campaignRepo) Create(campaign *models.Campaign) error {
	return r.db.Create(campaign).Error
}

func (r *campaignRepo) GetCampaignByName(campaignName string) (models.Campaign, error) {
	var campaign models.Campaign
	if err := r.db.First(&campaign).Where("campaign_name=?", campaignName).Error; err != nil {
		return campaign, errors.Wrapf(err, "failed to find campaign")
	}
	return campaign, nil
}
