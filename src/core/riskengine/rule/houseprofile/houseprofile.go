package houseprofile

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

const (
	mortgagedAdd = 1
)

func (r Rule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.HouseProfile.HouseStatus == entity.HsMortgaged {
		riskProfile.House.Value += mortgagedAdd
		riskProfile.Disability.Value += mortgagedAdd
	}

	return riskProfile
}
