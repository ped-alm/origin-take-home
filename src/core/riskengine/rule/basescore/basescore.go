package basescore

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
)

type Rule struct{}

func (r Rule) Execute(userProfile entity2.UserProfile, riskProfile entity2.RiskProfile) entity2.RiskProfile {
	var score = 0

	for _, question := range userProfile.RiskQuestions {
		if question.Answer == true {
			score++
		}
	}

	riskProfile.House.Value += score
	riskProfile.Auto.Value += score
	riskProfile.Disability.Value += score
	riskProfile.Life.Value += score

	return riskProfile
}
