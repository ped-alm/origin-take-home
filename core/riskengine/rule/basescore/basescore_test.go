package basescore

import (
	"github.com/ped-alm/origin-take-home/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRule_Execute(t *testing.T) {

	cases := []struct {
		expectedValue int
		answers       []bool
		name          string
	}{
		{
			name:          "case 0",
			expectedValue: 0,
			answers:       []bool{false, false, false}, // 0 0 0
		},
		{
			name:          "case 1",
			expectedValue: 1,
			answers:       []bool{false, false, true}, // 0 0 1
		},
		{
			name:          "case 2",
			expectedValue: 1,
			answers:       []bool{false, true, false}, // 0 1 0
		},
		{
			name:          "case 3",
			expectedValue: 2,
			answers:       []bool{false, true, true}, // 0 1 1
		},
		{
			name:          "case 4",
			expectedValue: 1,
			answers:       []bool{true, false, false}, // 1 0 0
		},
		{
			name:          "case 5",
			expectedValue: 2,
			answers:       []bool{true, false, true}, // 1 0 1
		},
		{
			name:          "case 6",
			expectedValue: 2,
			answers:       []bool{true, true, false}, // 1 1 0
		},
		{
			name:          "case 7",
			expectedValue: 3,
			answers:       []bool{true, true, true}, // 1 1 1
		},
	}

	for _, c := range cases {
		t.Run("should return the correct risk profile to all possible risk question answers - "+c.name, func(t *testing.T) {

			userProfile := entity.UserProfile{RiskQuestions: []entity.RiskQuestion{
				{
					RiskQuestionType: entity.RiskQuestion0,
					Answer:           c.answers[0],
				},
				{
					RiskQuestionType: entity.RiskQuestion1,
					Answer:           c.answers[1],
				},
				{
					RiskQuestionType: entity.RiskQuestion2,
					Answer:           c.answers[2],
				},
			}}

			received := Rule{}.Execute(userProfile, entity.RiskProfile{})

			expected := entity.RiskProfile{
				Disability: entity.Risk{Value: c.expectedValue},
				House:      entity.Risk{Value: c.expectedValue},
				Auto:       entity.Risk{Value: c.expectedValue},
				Life:       entity.Risk{Value: c.expectedValue},
			}

			assert.Equal(t, expected, received)
		})
	}
}
