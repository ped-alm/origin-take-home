package autoineligible

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user does not have income", func(t *testing.T) {
		userProfile := entity.UserProfile{
			Income: 0,
			VehicleProfile: entity.VehicleProfile{
				VehicleStatus: entity.VsOwned,
			},
			HouseProfile: entity.HouseProfile{
				HouseStatus: entity.HsOwned,
			},
		}

		received := Rule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Disability: entity.Risk{Status: entity.Ineligible},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user does not have a vehicle", func(t *testing.T) {
		userProfile := entity.UserProfile{
			Income: 1000,
			VehicleProfile: entity.VehicleProfile{
				VehicleStatus: entity.VsNotOwned,
			},
			HouseProfile: entity.HouseProfile{
				HouseStatus: entity.HsOwned,
			},
		}

		received := Rule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Auto: entity.Risk{Status: entity.Ineligible},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user does not have a house", func(t *testing.T) {
		userProfile := entity.UserProfile{
			Income: 1000,
			VehicleProfile: entity.VehicleProfile{
				VehicleStatus: entity.VsOwned,
			},
			HouseProfile: entity.HouseProfile{
				HouseStatus: entity.HsNotOwned,
			},
		}

		received := Rule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			House: entity.Risk{Status: entity.Ineligible},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user does not have an income, vehicle and house", func(t *testing.T) {
		userProfile := entity.UserProfile{
			Income: 0,
			VehicleProfile: entity.VehicleProfile{
				VehicleStatus: entity.VsNotOwned,
			},
			HouseProfile: entity.HouseProfile{
				HouseStatus: entity.HsNotOwned,
			},
		}

		received := Rule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Disability: entity.Risk{Status: entity.Ineligible},
			Auto:       entity.Risk{Status: entity.Ineligible},
			House:      entity.Risk{Status: entity.Ineligible},
		}

		assert.Equal(t, expected, received)
	})
}
