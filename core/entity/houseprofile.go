package entity

type HouseStatus int

const (
	HsNotOwned HouseStatus = iota
	HsOwned
	HsMortgaged
)

type HouseProfile struct {
	HouseStatus HouseStatus
}
