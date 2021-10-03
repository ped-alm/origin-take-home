package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type Engine struct {
	rules []RiskRule
}

func NewEngine() *Engine {
	var rules = []RiskRule{
		BaseScoreRule{},
		AutoIneligibleRule{},
		AgeRule{},
		IncomeRule{},
		HouseProfileRule{},
		DependentsRule{},
		MaritalRule{},
		VehicleProfileRule{},
	}

	return &Engine{rules}
}

func (e *Engine) Execute(userProfile entity.UserProfile) entity.RiskProfile {
	riskProfile := entity.RiskProfile{}

	for _, rule := range e.rules {
		riskProfile = rule.Execute(userProfile, riskProfile)
	}

	return riskProfile
}

func tweakRisk(risk entity.Risk, value int) entity.Risk {
	risk.Value += value

	switch {
	case risk.Status == entity.Ineligible:
		return risk
	case risk.Value <= 0:
		risk.Status = entity.Economic
	case risk.Value <= 2:
		risk.Status = entity.Regular
	default:
		risk.Status = entity.Responsible
	}

	return risk
}
