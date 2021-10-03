package entity

type RiskQuestionType int

// An identifier for which question is which was not given, so we are going to use just numbers
const (
	RiskQuestion0 RiskQuestionType = iota + 1
	RiskQuestion1
	RiskQuestion2
)

type RiskQuestion struct {
	RiskQuestionType RiskQuestionType
	Answer           bool
}
