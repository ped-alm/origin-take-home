package entity

type MaritalStatus int

const (
	Married MaritalStatus = iota
	Single
)

type UserProfile struct {
	Age            uint8
	Dependents     uint8
	HouseProfile   HouseProfile
	Income         uint64
	MaritalStatus  MaritalStatus
	RiskAnswers    []RiskQuestion
	VehicleProfile VehicleProfile
}
