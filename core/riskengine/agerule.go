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
		riskProfile.Disability.Status = entity.Ineligible
		riskProfile.Life.Status = entity.Ineligible
	case userProfile.Age < youngAge:
		riskProfile.House = tweakRisk(riskProfile.House, youngDeduct)
		riskProfile.Auto = tweakRisk(riskProfile.Auto, youngDeduct)
		riskProfile.Disability = tweakRisk(riskProfile.Disability, youngDeduct)
		riskProfile.Life = tweakRisk(riskProfile.Life, youngDeduct)
	case userProfile.Age < averageAge:
		riskProfile.House = tweakRisk(riskProfile.House, averageDeduct)
		riskProfile.Auto = tweakRisk(riskProfile.Auto, averageDeduct)
		riskProfile.Disability = tweakRisk(riskProfile.Disability, averageDeduct)
		riskProfile.Life = tweakRisk(riskProfile.Life, averageDeduct)
	}

	return riskProfile
}
