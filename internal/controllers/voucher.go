package controllers

import (
	"authen-system/internal/database"
	"github.com/pkg/errors"
)

type VoucherHandler interface {
	AllowUserUseVoucherInCampaign(req campaignRequest) error
}

type voucherHandler struct {
	voucherRepo  database.VoucherRepository
	campaignRepo database.CampaignRepository
}

func NewVoucherHandler(
	voucherRepo database.VoucherRepository,
	campaignRepo database.CampaignRepository,
) VoucherHandler {
	return &voucherHandler{
		voucherRepo:  voucherRepo,
		campaignRepo: campaignRepo,
	}
}

func (h *voucherHandler) AllowUserUseVoucherInCampaign(req campaignRequest) error {
	campaignInfo, err := h.campaignRepo.GetCampaignByName(req.campaignName)
	if err != nil {
		return errors.Wrapf(err, "failed to get campaign by name")
	}
	if campaignInfo.ID == 0 {
		return errors.New("campaign should be existed")
	}
	if err := h.voucherRepo.Create(database.VoucherRequest{
		CampaignID: campaignInfo.ID,
		UserID:     req.userID,
		EndDate:    campaignInfo.EndDate,
	}); err != nil {
		return err
	}
	return nil
}
