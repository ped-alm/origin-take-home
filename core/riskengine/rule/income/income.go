package income

import "github.com/ped-alm/origin-take-home/core/entity"

type Rule struct{}

const (
	highIncome       = 20000000 //$200.000,00
	highIncomeDeduct = -1
)

func (r Rule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {

	if userProfile.Income > highIncome {
		riskProfile.House.Value += highIncomeDeduct
		riskProfile.Auto.Value += highIncomeDeduct
		riskProfile.Disability.Value += highIncomeDeduct
		riskProfile.Life.Value += highIncomeDeduct
	}

	return riskProfile
}
