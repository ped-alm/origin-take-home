package model

type RiskProfile struct {
	Auto       string `json:"auto" binding:"required"`
	Disability string `json:"disability" binding:"required"`
	Home       string `json:"home" binding:"required"`
	Life       string `json:"life" binding:"required"`
}
