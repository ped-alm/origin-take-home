package dependents

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

const (
	dependentsAdd = 1
)

func (r Rule) Execute(userProfile entity2.UserProfile, riskProfile entity2.RiskProfile) entity2.RiskProfile {

	if userProfile.Dependents > 0 {
		riskProfile.Life.Value += dependentsAdd
		riskProfile.Disability.Value += dependentsAdd
	}

	return riskProfile
}
