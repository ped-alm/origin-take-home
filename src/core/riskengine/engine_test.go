package riskengine

import (
	"github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/age"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/autoineligible"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/basescore"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/dependents"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/houseprofile"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/income"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/marital"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/vehicleprofile"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEngine_Execute(t *testing.T) {

	cases := []struct {
		age           int
		dependents    int
		houseStatus   entity.HouseStatus
		vehicleYear   int
		income        int64
		maritalStatus entity.MaritalStatus
		answers       []bool
		name          string

		expectedDisability entity.RiskStatus
		expectedAuto       entity.RiskStatus
		expectedLife       entity.RiskStatus
		expectedHouse      entity.RiskStatus
	}{
		{
			age:           20,
			dependents:    0,
			houseStatus:   entity.HsNotOwned,
			vehicleYear:   0,
			income:        0,
			maritalStatus: entity.Single,
			answers:       []bool{false, false, false},
			name:          "case 0",

			expectedDisability: entity.Ineligible,
			expectedAuto:       entity.Ineligible,
			expectedLife:       entity.Economic,
			expectedHouse:      entity.Ineligible,
		},
		{
			age:           35,
			dependents:    2,
			houseStatus:   entity.HsOwned,
			vehicleYear:   time.Now().Year() - 10,
			income:        25000000,
			maritalStatus: entity.Married,
			answers:       []bool{false, false, true},
			name:          "case 1",

			expectedDisability: entity.Economic,
			expectedAuto:       entity.Economic,
			expectedLife:       entity.Regular,
			expectedHouse:      entity.Economic,
		},
		{
			age:           50,
			dependents:    0,
			houseStatus:   entity.HsMortgaged,
			vehicleYear:   time.Now().Year(),
			income:        17000000,
			maritalStatus: entity.Single,
			answers:       []bool{false, true, false},
			name:          "case 2",

			expectedDisability: entity.Regular,
			expectedAuto:       entity.Regular,
			expectedLife:       entity.Regular,
			expectedHouse:      entity.Regular,
		},
		{
			age:           71,
			dependents:    4,
			houseStatus:   entity.HsNotOwned,
			vehicleYear:   0,
			income:        0,
			maritalStatus: entity.Married,
			answers:       []bool{false, true, true},
			name:          "case 3",

			expectedDisability: entity.Ineligible,
			expectedAuto:       entity.Ineligible,
			expectedLife:       entity.Ineligible,
			expectedHouse:      entity.Ineligible,
		},
		{
			age:           27,
			dependents:    0,
			houseStatus:   entity.HsOwned,
			vehicleYear:   time.Now().Year() - 5,
			income:        21000000,
			maritalStatus: entity.Single,
			answers:       []bool{true, false, false},
			name:          "case 4",

			expectedDisability: entity.Economic,
			expectedAuto:       entity.Economic,
			expectedLife:       entity.Economic,
			expectedHouse:      entity.Economic,
		},
		{
			age:           31,
			dependents:    10,
			houseStatus:   entity.HsMortgaged,
			vehicleYear:   time.Now().Year() - 2,
			income:        12000000,
			maritalStatus: entity.Married,
			answers:       []bool{true, false, true},
			name:          "case 5",

			expectedDisability: entity.Regular,
			expectedAuto:       entity.Regular,
			expectedLife:       entity.Responsible,
			expectedHouse:      entity.Regular,
		},
		{
			age:           46,
			dependents:    0,
			houseStatus:   entity.HsNotOwned,
			vehicleYear:   0,
			income:        0,
			maritalStatus: entity.Single,
			answers:       []bool{true, true, false},
			name:          "case 6",

			expectedDisability: entity.Ineligible,
			expectedAuto:       entity.Ineligible,
			expectedLife:       entity.Regular,
			expectedHouse:      entity.Ineligible,
		},
		{
			age:           80,
			dependents:    3,
			houseStatus:   entity.HsOwned,
			vehicleYear:   time.Now().Year() - 7,
			income:        70000000,
			maritalStatus: entity.Married,
			answers:       []bool{true, true, true},
			name:          "case 7",

			expectedDisability: entity.Ineligible,
			expectedAuto:       entity.Regular,
			expectedLife:       entity.Ineligible,
			expectedHouse:      entity.Regular,
		},
	}

	for _, c := range cases {
		t.Run("should return the correct risk profile to all given user profiles "+c.name, func(t *testing.T) {
			var vehicleStatus entity.VehicleStatus
			if c.vehicleYear > 0 {
				vehicleStatus = entity.VsOwned
			}

			userProfile := entity.UserProfile{
				Age:        c.age,
				Dependents: c.dependents,
				HouseProfile: entity.HouseProfile{
					HouseStatus: c.houseStatus,
				},
				Income:        c.income,
				MaritalStatus: c.maritalStatus,
				RiskQuestions: []entity.RiskQuestion{
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
				},
				VehicleProfile: entity.VehicleProfile{
					VehicleStatus: vehicleStatus,
					Year:          c.vehicleYear,
				},
			}

			engine := NewEngine([]rule.Risk{
				basescore.Rule{},
				autoineligible.Rule{},
				age.Rule{},
				income.Rule{},
				houseprofile.Rule{},
				dependents.Rule{},
				marital.Rule{},
				vehicleprofile.Rule{},
			})

			received := engine.Execute(userProfile)

			assert.Equal(t, c.expectedLife, received.Life.Status)
			assert.Equal(t, c.expectedHouse, received.House.Status)
			assert.Equal(t, c.expectedDisability, received.Disability.Status)
			assert.Equal(t, c.expectedAuto, received.Auto.Status)
		})
	}
}
