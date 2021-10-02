package entity

type HouseStatus int

const (
	HsOwned HouseStatus = iota
	HsMortgaged
)

type HouseProfile struct {
	HouseStatus HouseStatus
}
