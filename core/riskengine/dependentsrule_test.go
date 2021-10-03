package riskengine

import (
	"github.com/ped-alm/origin-take-home/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDependentsRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user has dependents", func(t *testing.T) {
		userProfile := entity.UserProfile{Dependents: 2}
		received := DependentsRule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Disability: entity.Risk{Value: 1},
			Life:       entity.Risk{Value: 1},
		}

		assert.Equal(t, expected, received)
	})
}
