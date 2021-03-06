// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import (
	"time"

	"github.com/fleetdm/fleet/server/kolide"
)

var _ kolide.CampaignStore = (*CampaignStore)(nil)

type NewDistributedQueryCampaignFunc func(camp *kolide.DistributedQueryCampaign) (*kolide.DistributedQueryCampaign, error)

type DistributedQueryCampaignFunc func(id uint) (*kolide.DistributedQueryCampaign, error)

type SaveDistributedQueryCampaignFunc func(camp *kolide.DistributedQueryCampaign) error

type DistributedQueryCampaignTargetIDsFunc func(id uint) (hostIDs []uint, labelIDs []uint, err error)

type NewDistributedQueryCampaignTargetFunc func(target *kolide.DistributedQueryCampaignTarget) (*kolide.DistributedQueryCampaignTarget, error)

type CleanupDistributedQueryCampaignsFunc func(now time.Time) (expired uint, err error)

type CampaignStore struct {
	NewDistributedQueryCampaignFunc        NewDistributedQueryCampaignFunc
	NewDistributedQueryCampaignFuncInvoked bool

	DistributedQueryCampaignFunc        DistributedQueryCampaignFunc
	DistributedQueryCampaignFuncInvoked bool

	SaveDistributedQueryCampaignFunc        SaveDistributedQueryCampaignFunc
	SaveDistributedQueryCampaignFuncInvoked bool

	DistributedQueryCampaignTargetIDsFunc        DistributedQueryCampaignTargetIDsFunc
	DistributedQueryCampaignTargetIDsFuncInvoked bool

	NewDistributedQueryCampaignTargetFunc        NewDistributedQueryCampaignTargetFunc
	NewDistributedQueryCampaignTargetFuncInvoked bool

	CleanupDistributedQueryCampaignsFunc        CleanupDistributedQueryCampaignsFunc
	CleanupDistributedQueryCampaignsFuncInvoked bool
}

func (s *CampaignStore) NewDistributedQueryCampaign(camp *kolide.DistributedQueryCampaign) (*kolide.DistributedQueryCampaign, error) {
	s.NewDistributedQueryCampaignFuncInvoked = true
	return s.NewDistributedQueryCampaignFunc(camp)
}

func (s *CampaignStore) DistributedQueryCampaign(id uint) (*kolide.DistributedQueryCampaign, error) {
	s.DistributedQueryCampaignFuncInvoked = true
	return s.DistributedQueryCampaignFunc(id)
}

func (s *CampaignStore) SaveDistributedQueryCampaign(camp *kolide.DistributedQueryCampaign) error {
	s.SaveDistributedQueryCampaignFuncInvoked = true
	return s.SaveDistributedQueryCampaignFunc(camp)
}

func (s *CampaignStore) DistributedQueryCampaignTargetIDs(id uint) (hostIDs []uint, labelIDs []uint, err error) {
	s.DistributedQueryCampaignTargetIDsFuncInvoked = true
	return s.DistributedQueryCampaignTargetIDsFunc(id)
}

func (s *CampaignStore) NewDistributedQueryCampaignTarget(target *kolide.DistributedQueryCampaignTarget) (*kolide.DistributedQueryCampaignTarget, error) {
	s.NewDistributedQueryCampaignTargetFuncInvoked = true
	return s.NewDistributedQueryCampaignTargetFunc(target)
}

func (s *CampaignStore) CleanupDistributedQueryCampaigns(now time.Time) (expired uint, err error) {
	s.CleanupDistributedQueryCampaignsFuncInvoked = true
	return s.CleanupDistributedQueryCampaignsFunc(now)
}
