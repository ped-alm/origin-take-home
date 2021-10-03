package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type HouseProfileRule struct{}

const (
	mortgagedAdd = 1
)

func (r HouseProfileRule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.HouseProfile.HouseStatus == entity.HsMortgaged {
		riskProfile.House = tweakRisk(riskProfile.House, mortgagedAdd)
		riskProfile.Disability = tweakRisk(riskProfile.Disability, mortgagedAdd)
	}

	return riskProfile
}
