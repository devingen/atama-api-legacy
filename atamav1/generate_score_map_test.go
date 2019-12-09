package atamav1

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/devingen/atama-api/atama"
	"github.com/devingen/atama-api/model"
	"github.com/devingen/atama-api/testutil"
	"testing"
)

var generateScoreMapRules = []model.ConditionalComparisonRule{
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

type GenerateScoreMapTest struct {
	items1      []atama.MatchItem
	list1Fields []model.GenericField
	items2      []atama.MatchItem
	list2Fields []model.GenericField
	expected    ScoreMapMap
}

var generateScoreMapTests = []GenerateScoreMapTest{
	{
		items1: []atama.MatchItem{
			testutil.FirstItem10,
			testutil.FirstItem20,
			testutil.FirstItem21,
		},
		list1Fields: []model.GenericField{},
		items2: []atama.MatchItem{
			testutil.SecondItem10,
			testutil.SecondItem20,
			testutil.SecondItem30,
		},
		list2Fields: []model.GenericField{},
		expected: ScoreMapMap{
			testutil.FirstItem10["_id"].(string): {
				testutil.SecondItem30["_id"].(string): 1,
				testutil.SecondItem10["_id"].(string): 66,
			},
			testutil.FirstItem20["_id"].(string): {
				testutil.SecondItem20["_id"].(string): 1,
				testutil.SecondItem10["_id"].(string): 33,
				testutil.SecondItem30["_id"].(string): 33,
			},
			testutil.FirstItem21["_id"].(string): {
				testutil.SecondItem20["_id"].(string): 1,
				testutil.SecondItem10["_id"].(string): 33,
				testutil.SecondItem30["_id"].(string): 33,
			},
		},
	},
}

func TestGenerateScoreMap(t *testing.T) {
	for _, test := range generateScoreMapTests {
		result, _ := GenerateScoreMap(generateScoreMapRules, test.items1, test.items2, test.list1Fields, test.list2Fields)
		spew.Dump(result)

		for firstItemID, scoreMap := range result {
			for secondItemID, score := range scoreMap {
				expectedScore := test.expected[firstItemID][secondItemID]
				if expectedScore != score {
					t.Errorf("Case %v-%v: expected Score %v, got %v", firstItemID, secondItemID, expectedScore, score)
				}
			}
		}
	}
}
