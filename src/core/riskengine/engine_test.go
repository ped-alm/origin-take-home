package riskengine

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
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
		houseStatus   entity2.HouseStatus
		vehicleYear   int
		income        int64
		maritalStatus entity2.MaritalStatus
		answers       []bool
		name          string

		expectedDisability entity2.RiskStatus
		expectedAuto       entity2.RiskStatus
		expectedLife       entity2.RiskStatus
		expectedHouse      entity2.RiskStatus
	}{
		{
			age:           20,
			dependents:    0,
			houseStatus:   entity2.HsNotOwned,
			vehicleYear:   0,
			income:        0,
			maritalStatus: entity2.Single,
			answers:       []bool{false, false, false},
			name:          "case 0",

			expectedDisability: entity2.Ineligible,
			expectedAuto:       entity2.Ineligible,
			expectedLife:       entity2.Economic,
			expectedHouse:      entity2.Ineligible,
		},
		{
			age:           35,
			dependents:    2,
			houseStatus:   entity2.HsOwned,
			vehicleYear:   time.Now().Year() - 10,
			income:        25000000,
			maritalStatus: entity2.Married,
			answers:       []bool{false, false, true},
			name:          "case 1",

			expectedDisability: entity2.Economic,
			expectedAuto:       entity2.Economic,
			expectedLife:       entity2.Regular,
			expectedHouse:      entity2.Economic,
		},
		{
			age:           50,
			dependents:    0,
			houseStatus:   entity2.HsMortgaged,
			vehicleYear:   time.Now().Year(),
			income:        17000000,
			maritalStatus: entity2.Single,
			answers:       []bool{false, true, false},
			name:          "case 2",

			expectedDisability: entity2.Regular,
			expectedAuto:       entity2.Regular,
			expectedLife:       entity2.Regular,
			expectedHouse:      entity2.Regular,
		},
		{
			age:           71,
			dependents:    4,
			houseStatus:   entity2.HsNotOwned,
			vehicleYear:   0,
			income:        0,
			maritalStatus: entity2.Married,
			answers:       []bool{false, true, true},
			name:          "case 3",

			expectedDisability: entity2.Ineligible,
			expectedAuto:       entity2.Ineligible,
			expectedLife:       entity2.Ineligible,
			expectedHouse:      entity2.Ineligible,
		},
		{
			age:           27,
			dependents:    0,
			houseStatus:   entity2.HsOwned,
			vehicleYear:   time.Now().Year() - 5,
			income:        21000000,
			maritalStatus: entity2.Single,
			answers:       []bool{true, false, false},
			name:          "case 4",

			expectedDisability: entity2.Economic,
			expectedAuto:       entity2.Economic,
			expectedLife:       entity2.Economic,
			expectedHouse:      entity2.Economic,
		},
		{
			age:           31,
			dependents:    10,
			houseStatus:   entity2.HsMortgaged,
			vehicleYear:   time.Now().Year() - 2,
			income:        12000000,
			maritalStatus: entity2.Married,
			answers:       []bool{true, false, true},
			name:          "case 5",

			expectedDisability: entity2.Regular,
			expectedAuto:       entity2.Regular,
			expectedLife:       entity2.Responsible,
			expectedHouse:      entity2.Regular,
		},
		{
			age:           46,
			dependents:    0,
			houseStatus:   entity2.HsNotOwned,
			vehicleYear:   0,
			income:        0,
			maritalStatus: entity2.Single,
			answers:       []bool{true, true, false},
			name:          "case 6",

			expectedDisability: entity2.Ineligible,
			expectedAuto:       entity2.Ineligible,
			expectedLife:       entity2.Regular,
			expectedHouse:      entity2.Ineligible,
		},
		{
			age:           80,
			dependents:    3,
			houseStatus:   entity2.HsOwned,
			vehicleYear:   time.Now().Year() - 7,
			income:        70000000,
			maritalStatus: entity2.Married,
			answers:       []bool{true, true, true},
			name:          "case 7",

			expectedDisability: entity2.Ineligible,
			expectedAuto:       entity2.Regular,
			expectedLife:       entity2.Ineligible,
			expectedHouse:      entity2.Regular,
		},
	}

	for _, c := range cases {
		t.Run("should return the correct risk profile to all given user profiles "+c.name, func(t *testing.T) {
			var vehicleStatus entity2.VehicleStatus
			if c.vehicleYear > 0 {
				vehicleStatus = entity2.VsOwned
			}

			userProfile := entity2.UserProfile{
				Age:        c.age,
				Dependents: c.dependents,
				HouseProfile: entity2.HouseProfile{
					HouseStatus: c.houseStatus,
				},
				Income:        c.income,
				MaritalStatus: c.maritalStatus,
				RiskQuestions: []entity2.RiskQuestion{
					{
						RiskQuestionType: entity2.RiskQuestion0,
						Answer:           c.answers[0],
					},
					{
						RiskQuestionType: entity2.RiskQuestion1,
						Answer:           c.answers[1],
					},
					{
						RiskQuestionType: entity2.RiskQuestion2,
						Answer:           c.answers[2],
					},
				},
				VehicleProfile: entity2.VehicleProfile{
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
