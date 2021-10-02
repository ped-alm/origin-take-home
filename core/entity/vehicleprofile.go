package entity

type VehicleStatus int

const (
	VsOwned VehicleStatus = iota
	VsNotOwned
)

type VehicleProfile struct {
	VehicleStatus VehicleStatus
	Year          int
}
