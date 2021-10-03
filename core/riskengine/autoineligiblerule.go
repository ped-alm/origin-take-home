package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type AutoIneligibleRule struct{}

func (r AutoIneligibleRule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.Income == 0 {
		riskProfile.Disability.Status = entity.Ineligible
	}

	if userProfile.VehicleProfile.VehicleStatus == entity.VsNotOwned {
		riskProfile.Auto.Status = entity.Ineligible
	}

	if userProfile.HouseProfile.HouseStatus == entity.HsNotOwned {
		riskProfile.House.Status = entity.Ineligible
	}

	return riskProfile
}
