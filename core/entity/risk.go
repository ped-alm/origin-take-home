package entity

type RiskStatus int

const (
	Ineligible RiskStatus = iota + 1
	Economic
	Regular
	Responsible
)

type Risk struct {
	Value  int
	Status RiskStatus
}
