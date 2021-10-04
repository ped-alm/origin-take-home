package age

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

const (
	oldAge        = 60
	youngAge      = 30
	averageAge    = 40
	youngDeduct   = -2
	averageDeduct = -1
)

func (r Rule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	switch {
	case userProfile.Age > oldAge:
		riskProfile.Disability.Status = entity.Ineligible
		riskProfile.Life.Status = entity.Ineligible
	case userProfile.Age < youngAge:
		riskProfile.House.Value += youngDeduct
		riskProfile.Auto.Value += youngDeduct
		riskProfile.Disability.Value += youngDeduct
		riskProfile.Life.Value += youngDeduct
	case userProfile.Age < averageAge:
		riskProfile.House.Value += averageDeduct
		riskProfile.Auto.Value += averageDeduct
		riskProfile.Disability.Value += averageDeduct
		riskProfile.Life.Value += averageDeduct
	}

	return riskProfile
}
