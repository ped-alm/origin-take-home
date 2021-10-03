package marital

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

const (
	maritalAdd    = 1
	maritalDeduct = -1
)

func (r Rule) Execute(userProfile entity2.UserProfile, riskProfile entity2.RiskProfile) entity2.RiskProfile {

	if userProfile.MaritalStatus == entity2.Married {
		riskProfile.Life.Value += maritalAdd
		riskProfile.Disability.Value += maritalDeduct
	}

	return riskProfile
}
