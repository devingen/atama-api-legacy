package atama

import (
	"github.com/devingen/atama-api/model"
	"github.com/devingen/atama-api/testutil"
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
	items1      []MatchItem
	list1Fields []model.GenericField
	items2      []MatchItem
	list2Fields []model.GenericField
	expected    []MatchItemScores
}

var generateScoreMatrixTests = []GenerateScoreMatrixTest{
	{
		items1: []MatchItem{
			testutil.FirstItem10,
			testutil.FirstItem20,
			testutil.FirstItem21,
		},
		list1Fields: []model.GenericField{},
		items2: []MatchItem{
			testutil.SecondItem10,
			testutil.SecondItem20,
			testutil.SecondItem30,
		},
		list2Fields: []model.GenericField{},
		expected: []MatchItemScores{
			{
				Item: testutil.FirstItem10,
				Matches: []PairScore{
					{
						Item:  testutil.SecondItem30,
						Score: 1,
					},
					{
						Item:  testutil.SecondItem10,
						Score: 0.6666666666666666,
					},
					{
						Item:  testutil.SecondItem20,
						Score: 0,
					},
				},
			},
			{
				Item: testutil.FirstItem20,
				Matches: []PairScore{
					{
						Item:  testutil.SecondItem20,
						Score: 1,
					},
					{
						Item:  testutil.SecondItem10,
						Score: 0.3333333333333333,
					},
					{
						Item:  testutil.SecondItem30,
						Score: 0.3333333333333333,
					},
				},
			},
			{
				Item: testutil.FirstItem21,
				Matches: []PairScore{
					{
						Item:  testutil.SecondItem20,
						Score: 1,
					},
					{
						Item:  testutil.SecondItem10,
						Score: 0.3333333333333333,
					},
					{
						Item:  testutil.SecondItem30,
						Score: 0.3333333333333333,
					},
				},
			},
		},
	},
}

func TestGenerateScoreMatrix(t *testing.T) {
	for _, test := range generateScoreMatrixTests {
		result := GenerateScoreMatrix(generateScoreMatrixRules, test.items1, test.items2, test.list1Fields, test.list2Fields)

		for i, row := range result {
			for j, match := range row.Matches {
				expectedScore := test.expected[i].Matches[j].Score
				expectedID := test.expected[i].Matches[j].Item.GetID()

				if expectedID != match.Item.GetID() {
					t.Errorf("Case %d: expected ID %v, got %v", i, expectedID, match.Item.GetID())
				}
				if expectedScore != match.Score {
					t.Errorf("Case %d: expected Score %v, got %v", i, expectedScore, match.Score)
				}
			}
		}
	}
}
