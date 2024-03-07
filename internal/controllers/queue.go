package controllers

import (
	"authen-system/pkg/cache"
	"fmt"
)

type campaignRequest struct {
	campaignName string
	userID       uint
}
type CampaignQueue interface {
	Submit(req campaignRequest)
	Start()
}

const DefaultChannelSize = 10

type campaignQueue struct {
	queue          chan campaignRequest
	voucherHandler VoucherHandler
}

func NewCampaignQueue(
	voucherHandler VoucherHandler,
) CampaignQueue {
	return &campaignQueue{
		queue:          make(chan campaignRequest, DefaultChannelSize),
		voucherHandler: voucherHandler,
	}
}

func (h *campaignQueue) Submit(req campaignRequest) {
	h.queue <- req
}

func (h *campaignQueue) Start() {
	fmt.Println("waiting for jobs...")
	for {
		select {
		case request, ok := <-h.queue:
			if !ok {
				fmt.Println("channel has been closed")
				return
			}
			fmt.Printf("received job: %+v", request)
			if request.campaignName == cache.LoginFirstToTopupVoucher {
				if err := h.voucherHandler.AllowUserUseVoucherInCampaign(request); err != nil {
					fmt.Printf("error handling promotion: %v with %v", request, err)
				}
			} else {
				fmt.Printf("not handle: %v", request)
			}
		}
	}
}
