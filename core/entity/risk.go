package entity

type RiskStatus int

const (
	Ineligible RiskStatus = iota + 1
	Economic
	Regular
	Responsible
)

type Risk struct {
	value  int
	status RiskStatus
}

func (r *Risk) AddValue(value int) {
	r.value += value
	r.resetStatus()
}

func (r *Risk) resetStatus() {
	switch {
	case r.status == Ineligible:
		return
	case r.value <= 0:
		r.status = Economic
	case r.value <= 2:
		r.status = Regular
	default:
		r.status = Responsible
	}
}

func (r *Risk) Status() RiskStatus {
	return r.status
}

func (r *Risk) SetIneligible() {
	r.status = Ineligible
}
