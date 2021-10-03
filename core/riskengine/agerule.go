package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type AgeRule struct{}

const (
	oldAge        = 60
	youngAge      = 30
	averageAge    = 40
	youngDeduct   = -2
	averageDeduct = -1
)

func (r AgeRule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	switch {
	case userProfile.Age > oldAge:
		riskProfile.Disability.SetIneligible()
		riskProfile.Life.SetIneligible()
	case userProfile.Age < youngAge:
		riskProfile.House.AddValue(youngDeduct)
		riskProfile.Auto.AddValue(youngDeduct)
		riskProfile.Disability.AddValue(youngDeduct)
		riskProfile.Life.AddValue(youngDeduct)
	case userProfile.Age < averageAge:
		riskProfile.House.AddValue(averageDeduct)
		riskProfile.Auto.AddValue(averageDeduct)
		riskProfile.Disability.AddValue(averageDeduct)
		riskProfile.Life.AddValue(averageDeduct)
	}

	return riskProfile
}
