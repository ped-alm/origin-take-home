package entity

type RiskStatus int

const (
	Economic RiskStatus = iota
	Ineligible
	Regular
	Responsible
)

type Risk struct {
	Value  int
	Status RiskStatus
}
