package main

import (
	"github.com/ped-alm/origin-take-home/src/controller"
	"github.com/ped-alm/origin-take-home/src/core/riskengine"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/age"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/autoineligible"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/basescore"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/dependents"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/houseprofile"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/income"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/marital"
	"github.com/ped-alm/origin-take-home/src/core/riskengine/rule/vehicleprofile"
	"github.com/ped-alm/origin-take-home/src/infra/http"
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
