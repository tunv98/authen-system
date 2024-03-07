package cache

import (
	"sync"
)

const (
	LoginFirstToTopupVoucher = "TOPUP_FIRST_LOGIN"
)

type Cacheable interface {
	DecreaseCounter(campaignID string) bool
	GetCounter(campaignID string) int
	AddCampaigns(campaignID string, participants int)
}

var defaultCampaigns = map[string]int{
	LoginFirstToTopupVoucher: 100,
}

type cache struct {
	mutex           sync.Mutex
	campaignsCounts map[string]int
}

func NewCampaign() Cacheable {
	return &cache{
		campaignsCounts: defaultCampaigns,
	}
}

func (m *cache) DecreaseCounter(campaignID string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.campaignsCounts[campaignID] == 0 {
		delete(m.campaignsCounts, campaignID)
		return false
	}
	m.campaignsCounts[campaignID]--
	return true
}

func (m *cache) GetCounter(campaignID string) int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.campaignsCounts[campaignID]
}

func (m *cache) AddCampaigns(campaignID string, participants int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.campaignsCounts[campaignID] = participants
}
