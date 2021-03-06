package vehicleprofile

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user has a recent vehicle", func(t *testing.T) {
		userProfile := entity.UserProfile{VehicleProfile: entity.VehicleProfile{
			VehicleStatus: entity.VsOwned,
			Year:          time.Now().Year() - 1,
		}}
		received := Rule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Auto: entity.Risk{Value: 1},
		}

		assert.Equal(t, expected, received)
	})
}
