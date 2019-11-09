package model

type Comparison string

const (
	ComparisonEq        Comparison = "eq"
	ComparisonNe        Comparison = "ne"
	ComparisonLt        Comparison = "lt"
	ComparisonLte       Comparison = "lte"
	ComparisonGt        Comparison = "gt"
	ComparisonGte       Comparison = "gte"
	ComparisonIn        Comparison = "in"
	ComparisonContain   Comparison = "contain"
	ComparisonNcontain  Comparison = "ncontain"
	ComparisonSimilar   Comparison = "similar"
	ComparisonDifferent Comparison = "different"
)
