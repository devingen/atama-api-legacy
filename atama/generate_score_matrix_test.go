package atama

import (
	"github.com/devingen/atama-api/model"
	"testing"
)

var generateScoreMatrixConfig = model.CalculatorConfig{
	FieldID:    "email",
	FieldLimit: "limit",
}

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

type GenerateScoreMatrixTest struct {
	items1   []MatchItem
	items2   []MatchItem
	expected []MatchItemScores
}

var firstItem10 = MatchItem{
	"_id":        "can@doruk.com:0",
	"_v":         0,
	"email":      "can@doruk.com",
	"firstName":  "Can",
	"lastName":   "Doruk",
	"startYear":  2008,
	"department": "IT",
}

var firstItem20 = MatchItem{
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

var firstItem21 = MatchItem{
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

var secondItem10 = MatchItem{
	"_id":        "seda@candan.com",
	"_v":         0,
	"email":      "seda@candan.com",
	"firstName":  "Seda",
	"lastName":   "Candan",
	"startYear":  2010,
	"city":       "İstanbul",
	"department": "Sales",
}

var secondItem20 = MatchItem{
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

var secondItem30 = MatchItem{
	"_id":        "merve@toptan.com",
	"_v":         0,
	"email":      "merve@toptan.com",
	"firstName":  "Merve",
	"lastName":   "Toptan",
	"startYear":  2012,
	"department": "Sales",
	"limit":      2,
}

var generateScoreMatrixTests = []GenerateScoreMatrixTest{
	{
		items1: []MatchItem{
			firstItem10,
			firstItem20,
			firstItem21,
		},
		items2: []MatchItem{
			secondItem10,
			secondItem20,
			secondItem30,
		},
		expected: []MatchItemScores{
			{
				Item: firstItem10,
				Matches: []PairScore{
					{
						Item:  secondItem30,
						Score: 1,
					},
					{
						Item:  secondItem10,
						Score: 0.6666666666666666,
					},
					{
						Item:  secondItem20,
						Score: 0,
					},
				},
			},
			{
				Item: firstItem20,
				Matches: []PairScore{
					{
						Item:  secondItem20,
						Score: 1,
					},
					{
						Item:  secondItem10,
						Score: 0.3333333333333333,
					},
					{
						Item:  secondItem30,
						Score: 0.3333333333333333,
					},
				},
			},
			{
				Item: firstItem21,
				Matches: []PairScore{
					{
						Item:  secondItem20,
						Score: 1,
					},
					{
						Item:  secondItem10,
						Score: 0.3333333333333333,
					},
					{
						Item:  secondItem30,
						Score: 0.3333333333333333,
					},
				},
			},
		},
	},
}

func TestGenerateScoreMatrix(t *testing.T) {
	for _, test := range generateScoreMatrixTests {
		result := GenerateScoreMatrix(generateScoreMatrixConfig, generateScoreMatrixRules, test.items1, test.items2)

		for i, row := range result {
			for j, match := range row.Matches {
				expectedScore := test.expected[i].Matches[j].Score
				expectedID := test.expected[i].Matches[j].Item.GetID()

				if expectedID != match.Item.GetID() {
					t.Errorf("Case %d: expected ID %v, got %v", i, expectedID, match.Item.GetID())
				}
				if expectedScore != match.Score {
					t.Errorf("Case %d: expected ID %v, got %v", i, expectedScore, match.Score)
				}
			}
		}
	}
}
