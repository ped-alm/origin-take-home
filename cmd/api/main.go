package main

import (
	"github.com/ped-alm/origin-take-home/controller"
	"github.com/ped-alm/origin-take-home/core/riskengine"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule/age"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule/autoineligible"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule/basescore"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule/dependents"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule/houseprofile"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule/income"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule/marital"
	"github.com/ped-alm/origin-take-home/core/riskengine/rule/vehicleprofile"
	"github.com/ped-alm/origin-take-home/infra/http"
)

func main() {

	engine := riskengine.NewEngine([]rule.Risk{
		basescore.Rule{},
		autoineligible.Rule{},
		age.Rule{},
		income.Rule{},
		houseprofile.Rule{},
		dependents.Rule{},
		marital.Rule{},
		vehicleprofile.Rule{},
	})

	userController := controller.NewUser(*engine)

	gin := http.NewGin(*userController)

	gin.Start(":8080")
}
