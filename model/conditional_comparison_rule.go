package model

type ConditionalComparisonRule struct {
	Comparisons []RuleBase `json:"comparisons"`
	Type        RuleType   `json:"type"`
}
