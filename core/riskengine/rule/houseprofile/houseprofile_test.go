package houseprofile

import (
	"github.com/ped-alm/origin-take-home/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user house is mortgaged", func(t *testing.T) {
		userProfile := entity.UserProfile{HouseProfile: entity.HouseProfile{HouseStatus: entity.HsMortgaged}}
		received := Rule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Disability: entity.Risk{Value: 1},
			House:      entity.Risk{Value: 1},
		}

		assert.Equal(t, expected, received)
	})
}
