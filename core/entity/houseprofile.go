package entity

type HouseStatus int

const (
	HsOwned HouseStatus = iota + 1
	HsNotOwned
	HsMortgaged
)

type HouseProfile struct {
	HouseStatus HouseStatus
}
