package entity

type MaritalStatus int

const (
	Single MaritalStatus = iota
	Married
)

type UserProfile struct {
	Age            int
	Dependents     int
	HouseProfile   HouseProfile
	Income         int64
	MaritalStatus  MaritalStatus
	RiskQuestions  []RiskQuestion
	VehicleProfile VehicleProfile
}
