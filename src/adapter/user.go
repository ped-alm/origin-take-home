package adapter

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
	model2 "github.com/ped-alm/origin-take-home/src/infra/http/model"
)

func UserProfileHttpToEntity(mUserProfile model2.UserProfile) entity2.UserProfile {
	eUserProfile := entity2.UserProfile{
		Age:        mUserProfile.Age,
		Dependents: mUserProfile.Dependents,
		Income:     mUserProfile.Income,
	}

	switch mUserProfile.House.OwnershipStatus {
	case "owned":
		eUserProfile.HouseProfile.HouseStatus = entity2.HsOwned
	case "mortgaged":
		eUserProfile.HouseProfile.HouseStatus = entity2.HsMortgaged
	}

	switch mUserProfile.MaritalStatus {
	case "single":
		eUserProfile.MaritalStatus = entity2.Single
	case "married":
		eUserProfile.MaritalStatus = entity2.Married
	}

	for i, answer := range mUserProfile.RiskQuestions {
		var riskQuestion entity2.RiskQuestion

		switch i {
		case 0:
			riskQuestion.RiskQuestionType = entity2.RiskQuestion0
		case 1:
			riskQuestion.RiskQuestionType = entity2.RiskQuestion1
		case 2:
			riskQuestion.RiskQuestionType = entity2.RiskQuestion2
		}

		if answer == 0 {
			riskQuestion.Answer = false
		} else {
			riskQuestion.Answer = true
		}

		eUserProfile.RiskQuestions = append(eUserProfile.RiskQuestions, riskQuestion)
	}

	if mUserProfile.Vehicle.Year > 0 {
		eUserProfile.VehicleProfile.VehicleStatus = entity2.VsOwned
		eUserProfile.VehicleProfile.Year = mUserProfile.Vehicle.Year
	}

	return eUserProfile
}

func RiskProfileEntityToHttp(eRiskProfile entity2.RiskProfile) model2.RiskProfile {
	var mRiskProfile model2.RiskProfile

	mRiskProfile.Auto = riskStatusEntityToHttp(eRiskProfile.Auto.Status)
	mRiskProfile.Life = riskStatusEntityToHttp(eRiskProfile.Life.Status)
	mRiskProfile.Disability = riskStatusEntityToHttp(eRiskProfile.Disability.Status)
	mRiskProfile.Home = riskStatusEntityToHttp(eRiskProfile.House.Status)

	return mRiskProfile
}

func riskStatusEntityToHttp(status entity2.RiskStatus) string {
	switch status {
	case entity2.Regular:
		return "regular"
	case entity2.Responsible:
		return "responsible"
	case entity2.Ineligible:
		return "ineligible"
	case entity2.Economic:
		return "economic"
	default:
		return ""
	}
}
