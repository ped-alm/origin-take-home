package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type BaseScoreRule struct{}

func (r BaseScoreRule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {
	var score = 0

	for _, question := range userProfile.RiskQuestions {
		if question.Answer == true {
			score++
		}
	}

	riskProfile.House = tweakRisk(riskProfile.House, score)
	riskProfile.Auto = tweakRisk(riskProfile.Auto, score)
	riskProfile.Disability = tweakRisk(riskProfile.Disability, score)
	riskProfile.Life = tweakRisk(riskProfile.Life, score)

	return riskProfile
}
