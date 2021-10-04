package adapter

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/ped-alm/origin-take-home/src/infra/http/model"
)

func UserProfileHttpToEntity(mUserProfile model.UserProfile) entity.UserProfile {
	eUserProfile := entity.UserProfile{
		Age:        mUserProfile.Age,
		Dependents: mUserProfile.Dependents,
		Income:     mUserProfile.Income,
	}

	switch mUserProfile.House.OwnershipStatus {
	case "owned":
		eUserProfile.HouseProfile.HouseStatus = entity.HsOwned
	case "mortgaged":
		eUserProfile.HouseProfile.HouseStatus = entity.HsMortgaged
	}

	switch mUserProfile.MaritalStatus {
	case "single":
		eUserProfile.MaritalStatus = entity.Single
	case "married":
		eUserProfile.MaritalStatus = entity.Married
	}

	for i, answer := range mUserProfile.RiskQuestions {
		var riskQuestion entity.RiskQuestion

		switch i {
		case 0:
			riskQuestion.RiskQuestionType = entity.RiskQuestion0
		case 1:
			riskQuestion.RiskQuestionType = entity.RiskQuestion1
		case 2:
			riskQuestion.RiskQuestionType = entity.RiskQuestion2
		}

		if answer == 0 {
			riskQuestion.Answer = false
		} else {
			riskQuestion.Answer = true
		}

		eUserProfile.RiskQuestions = append(eUserProfile.RiskQuestions, riskQuestion)
	}

	if mUserProfile.Vehicle.Year > 0 {
		eUserProfile.VehicleProfile.VehicleStatus = entity.VsOwned
		eUserProfile.VehicleProfile.Year = mUserProfile.Vehicle.Year
	}

	return eUserProfile
}

func RiskProfileEntityToHttp(eRiskProfile entity.RiskProfile) model.RiskProfile {
	var mRiskProfile model.RiskProfile

	mRiskProfile.Auto = riskStatusEntityToHttp(eRiskProfile.Auto.Status)
	mRiskProfile.Life = riskStatusEntityToHttp(eRiskProfile.Life.Status)
	mRiskProfile.Disability = riskStatusEntityToHttp(eRiskProfile.Disability.Status)
	mRiskProfile.Home = riskStatusEntityToHttp(eRiskProfile.House.Status)

	return mRiskProfile
}

func riskStatusEntityToHttp(status entity.RiskStatus) string {
	switch status {
	case entity.Regular:
		return "regular"
	case entity.Responsible:
		return "responsible"
	case entity.Ineligible:
		return "ineligible"
	case entity.Economic:
		return "economic"
	default:
		return ""
	}
}
