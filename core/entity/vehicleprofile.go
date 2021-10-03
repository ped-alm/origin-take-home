package entity

type VehicleStatus int

const (
	VsOwned VehicleStatus = iota + 1
	VsNotOwned
)

type VehicleProfile struct {
	VehicleStatus VehicleStatus
	Year          int
}
