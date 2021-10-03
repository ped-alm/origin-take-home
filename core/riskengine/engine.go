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
