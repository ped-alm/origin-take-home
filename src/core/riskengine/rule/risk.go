package rule

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
)

type Risk interface {
	Execute(userProfile entity2.UserProfile, riskProfile entity2.RiskProfile) entity2.RiskProfile
}
