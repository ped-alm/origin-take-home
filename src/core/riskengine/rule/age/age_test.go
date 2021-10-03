package age

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user is above old age", func(t *testing.T) {
		userProfile := entity2.UserProfile{Age: 61}
		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			Disability: entity2.Risk{Status: entity2.Ineligible},
			Life:       entity2.Risk{Status: entity2.Ineligible},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user is under young age", func(t *testing.T) {
		userProfile := entity2.UserProfile{Age: 29}
		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			Auto:       entity2.Risk{Value: -2},
			Disability: entity2.Risk{Value: -2},
			House:      entity2.Risk{Value: -2},
			Life:       entity2.Risk{Value: -2},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user is under average age and above young age", func(t *testing.T) {
		userProfile := entity2.UserProfile{Age: 30}
		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})

		expected := entity2.RiskProfile{
			Auto:       entity2.Risk{Value: -1},
			Disability: entity2.Risk{Value: -1},
			House:      entity2.Risk{Value: -1},
			Life:       entity2.Risk{Value: -1},
		}

		assert.Equal(t, expected, received)
	})

	t.Run("should return the correct risk profile when the user is above average age and under old age", func(t *testing.T) {
		userProfile := entity2.UserProfile{Age: 40}
		received := Rule{}.Execute(userProfile, entity2.RiskProfile{})
		expected := entity2.RiskProfile{}

		assert.Equal(t, expected, received)
	})
}
