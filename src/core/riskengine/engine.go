package riskengine

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule"
)

type Engine struct {
	rules []rule.Risk
}

func NewEngine(rules []rule.Risk) *Engine {
	return &Engine{rules}
}

func (e *Engine) Execute(userProfile entity2.UserProfile) entity2.RiskProfile {
	riskProfile := entity2.RiskProfile{}

	for _, r := range e.rules {
		riskProfile = r.Execute(userProfile, riskProfile)

		riskProfile.Disability = tweakRisk(riskProfile.Disability)
		riskProfile.Auto = tweakRisk(riskProfile.Auto)
		riskProfile.Life = tweakRisk(riskProfile.Life)
		riskProfile.House = tweakRisk(riskProfile.House)
	}

	return riskProfile
}

func tweakRisk(risk entity2.Risk) entity2.Risk {

	switch {
	case risk.Status == entity2.Ineligible:
		return risk
	case risk.Value <= 0:
		risk.Status = entity2.Economic
	case risk.Value <= 2:
		risk.Status = entity2.Regular
	default:
		risk.Status = entity2.Responsible
	}

	return risk
}
