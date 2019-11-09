package model

type RuleBase struct {
	FirstField  *BaseField
	FirstValue  interface{}
	Comparison  Comparison
	SecondField *BaseField
	SecondValue interface{}
}
