package houseprofile

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user house is mortgaged", func(t *testing.T) {
		userProfile := entity2.UserProfile{HouseProfile: entity2.HouseProfile{HouseStatus: entity2.HsMortgaged}}
		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			Disability: entity2.Risk{Value: 1},
			House:      entity2.Risk{Value: 1},
		}

		assert.Equal(t, expected, received)
	})
}
