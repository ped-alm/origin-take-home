package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type MaritalRule struct{}

const (
	maritalAdd    = 1
	maritalDeduct = -1
)

func (r MaritalRule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.MaritalStatus == entity.Married {
		riskProfile.Life.Value += maritalAdd
		riskProfile.Disability.Value += maritalDeduct
	}

	return riskProfile
}
