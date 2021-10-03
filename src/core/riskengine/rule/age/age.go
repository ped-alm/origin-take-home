package age

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

const (
	oldAge        = 60
	youngAge      = 30
	averageAge    = 40
	youngDeduct   = -2
	averageDeduct = -1
)

func (r Rule) Execute(userProfile entity2.UserProfile, riskProfile entity2.RiskProfile) entity2.RiskProfile {

	switch {
	case userProfile.Age > oldAge:
		riskProfile.Disability.Status = entity2.Ineligible
		riskProfile.Life.Status = entity2.Ineligible
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
