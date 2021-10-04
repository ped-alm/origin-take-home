package vehicleprofile

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
	"time"
)

type Rule struct{}

const (
	newVehicleAdd       = 1
	newVehicleThreshold = 5
)

func (r Rule) Execute(userProfile entity.UserProfile, riskProfile entity.RiskProfile) entity.RiskProfile {
	newThresholdYear := time.Now().Year() - newVehicleThreshold

	if userProfile.VehicleProfile.VehicleStatus == entity.VsOwned && userProfile.VehicleProfile.Year >= newThresholdYear {
		riskProfile.Auto.Value += newVehicleAdd
	}

	return riskProfile
}
