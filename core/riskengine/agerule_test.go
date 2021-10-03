package riskengine

import (
	"github.com/ped-alm/origin-take-home/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAgeRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user is above old age", func(t *testing.T) {
		userProfile := entity.UserProfile{Age: 61}
		received := AgeRule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Disability: entity.Risk{Status: entity.Ineligible},
			Life:       entity.Risk{Status: entity.Ineligible},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user is under young age", func(t *testing.T) {
		userProfile := entity.UserProfile{Age: 29}
		received := AgeRule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Auto:       entity.Risk{Value: -2},
			Disability: entity.Risk{Value: -2},
			House:      entity.Risk{Value: -2},
			Life:       entity.Risk{Value: -2},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user is under average age and above young age", func(t *testing.T) {
		userProfile := entity.UserProfile{Age: 30}
		received := AgeRule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Auto:       entity.Risk{Value: -1},
			Disability: entity.Risk{Value: -1},
			House:      entity.Risk{Value: -1},
			Life:       entity.Risk{Value: -1},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user is above average age and under old age", func(t *testing.T) {
		userProfile := entity.UserProfile{Age: 40}
		received := AgeRule{}.Execute(userProfile, entity.RiskProfile{})
		expected := entity.RiskProfile{}

		assert.Equal(t, expected, received)
	})
}
