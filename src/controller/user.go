package controller

import (
	entity2 "github.com/ped-alm/origin-take-home/src/core/entity"
	"github.com/ped-alm/origin-take-home/src/core/riskengine"
)

type User struct {
	engine riskengine.Engine
}

func NewUser(engine riskengine.Engine) *User {
	return &User{engine}
}

func (u *User) CalculateRisk(userProfile entity2.UserProfile) entity2.RiskProfile {
	return u.engine.Execute(userProfile)
}
