package entity

type MaritalStatus int

const (
	Married MaritalStatus = iota + 1
	Single
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
