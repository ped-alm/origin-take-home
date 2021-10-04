package dependents

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

const (
	dependentsAdd = 1
)

func (r Rule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.Dependents > 0 {
		riskProfile.Life.Value += dependentsAdd
		riskProfile.Disability.Value += dependentsAdd
	}

	return riskProfile
}
