package autoineligible

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

func (r Rule) Execute(userProfile entity2.UserProfile, riskProfile entity2.RiskProfile) entity2.RiskProfile {

	if userProfile.Income == 0 {
		riskProfile.Disability.Status = entity2.Ineligible
	}

	if userProfile.VehicleProfile.VehicleStatus == entity2.VsNotOwned {
		riskProfile.Auto.Status = entity2.Ineligible
	}

	if userProfile.HouseProfile.HouseStatus == entity2.HsNotOwned {
		riskProfile.House.Status = entity2.Ineligible
	}

	return riskProfile
}
