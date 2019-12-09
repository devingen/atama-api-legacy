package testutil

import (
	"github.com/devingen/atama-api/model"
)

var generateScoreMatrixRules = []model.ConditionalComparisonRule{
	{
		Comparisons: []model.RuleBase{
			{
				FirstField:  &model.BaseField{ID: "city"},
				FirstValue:  nil,
				Comparison:  model.ComparisonEq,
				SecondField: &model.BaseField{ID: "city"},
				SecondValue: nil,
			},
		},
		Type: model.RuleTypeComparison,
	},
	{
		Comparisons: []model.RuleBase{
			{
				FirstField:  &model.BaseField{ID: "department"},
				FirstValue:  nil,
				Comparison:  model.ComparisonNe,
				SecondField: &model.BaseField{ID: "department"},
				SecondValue: nil,
			},
		},
		Type: model.RuleTypeComparison,
	},
	{
		Comparisons: []model.RuleBase{
			{
				FirstField:  &model.BaseField{ID: "startYear"},
				FirstValue:  nil,
				Comparison:  model.ComparisonLt,
				SecondField: &model.BaseField{ID: "startYear"},
				SecondValue: nil,
			},
		},
		Type: model.RuleTypeComparison,
	},
}

var FirstItem10 = map[string]interface{}{
	"_id":        "can@doruk.com:0",
	"_v":         0,
	"email":      "can@doruk.com",
	"firstName":  "Can",
	"lastName":   "Doruk",
	"startYear":  2008,
	"department": "IT",
}

var FirstItem20 = map[string]interface{}{
	"_id":        "kamil@boyer.com:0",
	"_v":         0,
	"email":      "kamil@boyer.com",
	"firstName":  "Kamil",
	"lastName":   "Boyer",
	"startYear":  2006,
	"city":       "Ankara",
	"department": "Sales",
	"limit":      2,
}

var FirstItem21 = map[string]interface{}{
	"_id":        "kamil@boyer.com:1",
	"_v":         1,
	"email":      "kamil@boyer.com",
	"firstName":  "Kamil",
	"lastName":   "Boyer",
	"startYear":  2006,
	"city":       "Ankara",
	"department": "Sales",
	"limit":      2,
}

var SecondItem10 = map[string]interface{}{
	"_id":        "seda@candan.com",
	"_v":         0,
	"email":      "seda@candan.com",
	"firstName":  "Seda",
	"lastName":   "Candan",
	"startYear":  2010,
	"city":       "İstanbul",
	"department": "Sales",
}

var SecondItem20 = map[string]interface{}{
	"_id":        "mert@kudret.com",
	"_v":         0,
	"email":      "mert@kudret.com",
	"firstName":  "Mert",
	"lastName":   "Kudret",
	"startYear":  2007,
	"city":       "Ankara",
	"department": "IT",
	"limit":      2,
}

var SecondItem30 = map[string]interface{}{
	"_id":        "merve@toptan.com",
	"_v":         0,
	"email":      "merve@toptan.com",
	"firstName":  "Merve",
	"lastName":   "Toptan",
	"startYear":  2012,
	"department": "Sales",
	"limit":      2,
}
