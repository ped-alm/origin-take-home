package riskengine

import "github.com/ped-alm/origin-take-home/core/entity"

type RiskRule interface {
	Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile
}
