package marital

import (
	"github.com/ped-alm/origin-take-home/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Execute(t *testing.T) {

	t.Run("should return the correct risk profile when the user is married", func(t *testing.T) {
		userProfile := entity.UserProfile{MaritalStatus: entity.Married}
		received := Rule{}.Execute(userProfile, entity.RiskProfile{})

		expected := entity.RiskProfile{
			Disability: entity.Risk{Value: -1},
			Life:       entity.Risk{Value: 1},
		}

		assert.Equal(t, expected, received)
	})
}
