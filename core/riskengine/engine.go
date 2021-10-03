package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type Engine struct {
	rules []RiskRule
}

func NewEngine(rules []RiskRule) *Engine {
	return &Engine{rules}
}

func (e *Engine) Execute(userProfile entity.UserProfile) entity.RiskProfile {
	riskProfile := entity.RiskProfile{}

	for _, rule := range e.rules {
		riskProfile = rule.Execute(userProfile, riskProfile)

		riskProfile.Disability = tweakRisk(riskProfile.Disability)
		riskProfile.Auto = tweakRisk(riskProfile.Auto)
		riskProfile.Life = tweakRisk(riskProfile.Life)
		riskProfile.House = tweakRisk(riskProfile.House)
	}

	return riskProfile
}

func tweakRisk(risk entity.Risk) entity.Risk {

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
