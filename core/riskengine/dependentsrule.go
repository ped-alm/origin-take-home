package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type DependentsRule struct{}

const (
	dependentsAdd = 1
)

func (r DependentsRule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.Dependents > 0 {
		riskProfile.Life.AddValue(dependentsAdd)
		riskProfile.Disability.AddValue(dependentsAdd)
	}

	return riskProfile
}
