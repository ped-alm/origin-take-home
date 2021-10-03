package houseprofile

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

const (
	mortgagedAdd = 1
)

func (r Rule) Execute(userProfile entity2.UserProfile, riskProfile entity2.RiskProfile) entity2.RiskProfile {

	if userProfile.HouseProfile.HouseStatus == entity2.HsMortgaged {
		riskProfile.House.Value += mortgagedAdd
		riskProfile.Disability.Value += mortgagedAdd
	}

	return riskProfile
}
