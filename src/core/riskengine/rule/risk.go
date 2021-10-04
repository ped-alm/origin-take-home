package rule

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
)

type Risk interface {
	Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile
}
