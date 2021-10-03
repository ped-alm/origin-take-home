package vehicleprofile

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
	"time"
)

type Rule struct{}

const (
	newVehicleAdd       = 1
	newVehicleThreshold = 5
)

func (r Rule) Execute(userProfile entity2.UserProfile, riskProfile entity2.RiskProfile) entity2.RiskProfile {
	newThresholdYear := time.Now().Year() - newVehicleThreshold

	if userProfile.VehicleProfile.VehicleStatus == entity2.VsOwned && userProfile.VehicleProfile.Year >= newThresholdYear {
		riskProfile.Auto.Value += newVehicleAdd
	}

	return riskProfile
}
