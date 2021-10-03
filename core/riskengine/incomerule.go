package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type IncomeRule struct{}

const (
	highIncome       = 20000000 //$200.000,00
	highIncomeDeduct = -1
)

func (r IncomeRule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.Income > highIncome {
		riskProfile.House = tweakRisk(riskProfile.House, highIncomeDeduct)
		riskProfile.Auto = tweakRisk(riskProfile.Auto, highIncomeDeduct)
		riskProfile.Disability = tweakRisk(riskProfile.Disability, highIncomeDeduct)
		riskProfile.Life = tweakRisk(riskProfile.Life, highIncomeDeduct)
	}

	return riskProfile
}
