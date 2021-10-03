package vehicleprofile

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user has a recent vehicle", func(t *testing.T) {
		userProfile := entity2.UserProfile{VehicleProfile: entity2.VehicleProfile{
			VehicleStatus: entity2.VsOwned,
			Year:          time.Now().Year() - 1,
		}}
		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			Auto: entity2.Risk{Value: 1},
		}

		assert.Equal(t, expected, received)
	})
}
