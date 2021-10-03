package model

type UserProfile struct {
	Age           int     `json:"age"`
	Dependents    int     `json:"dependents"`
	House         House   `json:"house"`
	Income        int64   `json:"income"`
	MaritalStatus string  `json:"marital_status" binding:"required"`
	RiskQuestions []int   `json:"risk_questions" binding:"required"`
	Vehicle       Vehicle `json:"vehicle" binding:"required"`
}

type House struct {
	OwnershipStatus string `json:"ownership_status"`
}

type Vehicle struct {
	Year int `json:"year"`
}
