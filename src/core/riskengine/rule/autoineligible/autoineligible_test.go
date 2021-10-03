package autoineligible

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user does not have income", func(t *testing.T) {
		userProfile := entity2.UserProfile{
			Income: 0,
			VehicleProfile: entity2.VehicleProfile{
				VehicleStatus: entity2.VsOwned,
			},
			HouseProfile: entity2.HouseProfile{
				HouseStatus: entity2.HsOwned,
			},
		}

		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			Disability: entity2.Risk{Status: entity2.Ineligible},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user does not have a vehicle", func(t *testing.T) {
		userProfile := entity2.UserProfile{
			Income: 1000,
			VehicleProfile: entity2.VehicleProfile{
				VehicleStatus: entity2.VsNotOwned,
			},
			HouseProfile: entity2.HouseProfile{
				HouseStatus: entity2.HsOwned,
			},
		}

		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			Auto: entity2.Risk{Status: entity2.Ineligible},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user does not have a house", func(t *testing.T) {
		userProfile := entity2.UserProfile{
			Income: 1000,
			VehicleProfile: entity2.VehicleProfile{
				VehicleStatus: entity2.VsOwned,
			},
			HouseProfile: entity2.HouseProfile{
				HouseStatus: entity2.HsNotOwned,
			},
		}

		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			House: entity2.Risk{Status: entity2.Ineligible},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user does not have an income, vehicle and house", func(t *testing.T) {
		userProfile := entity2.UserProfile{
			Income: 0,
			VehicleProfile: entity2.VehicleProfile{
				VehicleStatus: entity2.VsNotOwned,
			},
			HouseProfile: entity2.HouseProfile{
				HouseStatus: entity2.HsNotOwned,
			},
		}

		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			Disability: entity2.Risk{Status: entity2.Ineligible},
			Auto:       entity2.Risk{Status: entity2.Ineligible},
			House:      entity2.Risk{Status: entity2.Ineligible},
		}

		assert.Equal(t, expected, received)
	})
}
