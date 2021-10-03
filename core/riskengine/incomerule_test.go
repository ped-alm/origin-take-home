package riskengine

import (
	"github.com/ped-alm/origin-take-home/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncomeRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user has a high income", func(t *testing.T) {
		userProfile := entity.UserProfile{Income: 20100000} //  $201.000,00
		received := IncomeRule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Disability: entity.Risk{Value: -1},
			House:      entity.Risk{Value: -1},
			Life:       entity.Risk{Value: -1},
			Auto:       entity.Risk{Value: -1},
		}

		assert.Equal(t, expected, received)
	})
}