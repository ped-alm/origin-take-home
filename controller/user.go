package controller

import (
	"github.com/ped-alm/origin-take-home/core/entity"
	"github.com/ped-alm/origin-take-home/core/riskengine"
)

type User struct {
	engine riskengine.Engine
}

func NewUser(engine riskengine.Engine) *User {
	return &User{engine}
}

func (u *User) CalculateRisk(userProfile entity.UserProfile) entity.RiskProfile {
	return u.engine.Execute(userProfile)
}
