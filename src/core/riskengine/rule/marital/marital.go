package marital

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

const (
	maritalAdd    = 1
	maritalDeduct = -1
)

func (r Rule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.MaritalStatus == entity.Married {
		riskProfile.Life.Value += maritalAdd
		riskProfile.Disability.Value += maritalDeduct
	}

	return riskProfile
}
