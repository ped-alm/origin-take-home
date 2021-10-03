package basescore

import "github.com/ped-alm/origin-take-home/core/entity"

type Rule struct{}

func (r Rule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {
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
