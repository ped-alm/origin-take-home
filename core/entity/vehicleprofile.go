package entity

type VehicleStatus int

const (
	VsNotOwned VehicleStatus = iota
	VsOwned
)

type VehicleProfile struct {
	VehicleStatus VehicleStatus
	Year          int
}
