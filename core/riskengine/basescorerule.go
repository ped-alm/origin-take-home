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

	riskProfile.House.AddValue(score)
	riskProfile.Auto.AddValue(score)
	riskProfile.Disability.AddValue(score)
	riskProfile.Life.AddValue(score)

	return riskProfile
}
