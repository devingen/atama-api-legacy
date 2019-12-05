package atama

import (
	"github.com/devingen/atama-api/model"
	"testing"
)

// region TestCalculateSimilarity

type CalculateSimilarityTest struct {
	array1   []interface{}
	array2   []interface{}
	expected float64
}

var calculateSimilarityTests = []CalculateSimilarityTest{
	{[]interface{}{1, 2}, []interface{}{1, 2}, 1},
	{[]interface{}{1, 2}, []interface{}{2, 3}, 0.3333333333333333},
	{[]interface{}{1, 2}, []interface{}{3}, 0},
}

func TestCalculateSimilarity(t *testing.T) {
	for i, test := range calculateSimilarityTests {
		result := calculateSimilarity(test.array1, test.array2)
		if result != test.expected {
			t.Errorf("Case %d: expected %v, got %v", i, test.expected, result)
		}
	}
}

// endregion

// region TestCompareValues

type CompareValuesTest struct {
	comparison model.Comparison
	value1     interface{}
	value2     interface{}
	expected   float64
}

var compareValuesTests = []CompareValuesTest{
	{model.ComparisonEq, 1, 1, 1},
	{model.ComparisonEq, 1, 2, 0},
	{model.ComparisonEq, "a", "a", 1},
	{model.ComparisonEq, "a", "b", 0},
	{model.ComparisonEq, "a", 1, 0},
	{model.ComparisonNe, 1, 1, 0},
	{model.ComparisonNe, 1, 2, 1},
	{model.ComparisonNe, "a", "a", 0},
	{model.ComparisonNe, "a", "b", 1},
	{model.ComparisonNe, "a", 1, 1},
	{model.ComparisonLt, 1, 0, 0},
	{model.ComparisonLt, 1, 1, 0},
	{model.ComparisonLt, 1, 2, 1},
	{model.ComparisonLt, "a", 2, 1},
	{model.ComparisonLte, 1, 0, 0},
	{model.ComparisonLte, 1, 1, 1},
	{model.ComparisonLte, 1, 2, 1},
	{model.ComparisonLte, "a", 2, 1},
	{model.ComparisonGt, 1, 0, 1},
	{model.ComparisonGt, 1, 1, 0},
	{model.ComparisonGt, 1, 2, 0},
	{model.ComparisonGt, "a", 2, 0},
	{model.ComparisonGte, 1, 0, 1},
	{model.ComparisonGte, 1, 1, 1},
	{model.ComparisonGte, 1, 2, 0},
	{model.ComparisonGte, "a", 2, 0},
	{model.ComparisonIn, "a", []interface{}{"a", "b"}, 1},
	{model.ComparisonIn, "c", []interface{}{"a", "b"}, 0},
	{model.ComparisonIn, 1, []interface{}{"a", "b"}, 0},
	{model.ComparisonSimilar, []interface{}{"a", "b"}, []interface{}{"a", "b"}, 1},
	{model.ComparisonSimilar, []interface{}{"a"}, []interface{}{"a", "b"}, 0.5},
	{model.ComparisonSimilar, []interface{}{"a", "b", "c"}, []interface{}{"a", "b"}, 0.6666666666666666},
	{model.ComparisonDifferent, []interface{}{"a", "b"}, []interface{}{"a", "b"}, 0},
	{model.ComparisonDifferent, []interface{}{"a"}, []interface{}{"a", "b"}, 0.5},
	{model.ComparisonDifferent, []interface{}{"a", "b", "c"}, []interface{}{"a", "b"}, 0.33333333333333337},
	{model.ComparisonContain, []interface{}{"a", "b"}, []interface{}{"a"}, 1},
	{model.ComparisonContain, []interface{}{"a", "b"}, []interface{}{"a", "b"}, 1},
	{model.ComparisonContain, []interface{}{"a", "b"}, []interface{}{"a", "b", "c"}, 0},
	{model.ComparisonNcontain, []interface{}{"a", "b"}, []interface{}{"a"}, 0},
	{model.ComparisonNcontain, []interface{}{"a", "b"}, []interface{}{"a", "b"}, 0},
	{model.ComparisonNcontain, []interface{}{"a", "b"}, []interface{}{"a", "b", "c"}, 1},
}

func TestCompareValues(t *testing.T) {
	for i, test := range compareValuesTests {
		result := compareValues(test.comparison, test.value1, test.value2)
		if result != test.expected {
			t.Errorf("Case %d: expected %v, got %v", i, test.expected, result)
		}
	}
}

// endregion

// region CalculatePair

type CalculatePairTest struct {
	rules                    []model.ConditionalComparisonRule
	data1                    map[string]interface{}
	list1FieldOptionLabelMap FieldOptionLabelMap
	data2                    map[string]interface{}
	list2FieldOptionLabelMap FieldOptionLabelMap
	expected                 float64
}

var rules1 = []model.ConditionalComparisonRule{
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
}

var rules2 = []model.ConditionalComparisonRule{
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
				FirstField:  &model.BaseField{ID: "birthYear"},
				FirstValue:  nil,
				Comparison:  model.ComparisonNe,
				SecondField: &model.BaseField{ID: "birthYear"},
				SecondValue: nil,
			},
		},
		Type: model.RuleTypeComparison,
	},
}

var rules3 = []model.ConditionalComparisonRule{
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
				FirstField:  &model.BaseField{ID: "birthYear"},
				FirstValue:  nil,
				Comparison:  model.ComparisonNe,
				SecondField: &model.BaseField{ID: "birthYear"},
				SecondValue: nil,
			},
		},
		Type: model.RuleTypeComparison,
	},
	{
		Comparisons: []model.RuleBase{
			{
				FirstField:  &model.BaseField{ID: "skills"},
				FirstValue:  nil,
				Comparison:  model.ComparisonSimilar,
				SecondField: &model.BaseField{ID: "skills"},
				SecondValue: nil,
			},
		},
		Type: model.RuleTypeComparison,
	},
}

var rules4 = []model.ConditionalComparisonRule{
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
				FirstField:  &model.BaseField{ID: "birthYear"},
				FirstValue:  nil,
				Comparison:  model.ComparisonNe,
				SecondField: &model.BaseField{ID: "birthYear"},
				SecondValue: nil,
			},
		},
		Type: model.RuleTypeComparison,
	},
	{
		Comparisons: []model.RuleBase{
			{
				FirstField:  &model.BaseField{ID: "skills"},
				FirstValue:  nil,
				Comparison:  model.ComparisonSimilar,
				SecondField: &model.BaseField{ID: "skills"},
				SecondValue: nil,
			},
		},
		Type: model.RuleTypeComparison,
	},
	{
		Comparisons: []model.RuleBase{
			{ // if the city of the first data equals to İstanbul
				FirstField:  &model.BaseField{ID: "city"},
				FirstValue:  nil,
				Comparison:  model.ComparisonEq,
				SecondField: nil,
				SecondValue: "İstanbul",
			},
			{ // then the city of the second data should be Edirne
				FirstField:  &model.BaseField{ID: "city"},
				FirstValue:  nil,
				Comparison:  model.ComparisonEq,
				SecondField: nil,
				SecondValue: "Edirne",
			},
		},
		Type: model.RuleTypeConditionalComparison,
	},
}

var calculatePairTests = []CalculatePairTest{
	{
		rules1,
		map[string]interface{}{
			"city": "İstanbul",
		},
		map[string]map[interface{}]interface{}{},
		map[string]interface{}{
			"city": "İstanbul",
		},
		map[string]map[interface{}]interface{}{},
		1,
	},
	{
		rules1,
		map[string]interface{}{
			"city": "Edirne",
		},
		map[string]map[interface{}]interface{}{},
		map[string]interface{}{
			"city": "İstanbul",
		},
		map[string]map[interface{}]interface{}{},
		0,
	},
	{
		rules2,
		map[string]interface{}{
			"city":      "Edirne",
			"birthYear": 1432,
		},
		map[string]map[interface{}]interface{}{},
		map[string]interface{}{
			"city":      "İstanbul",
			"birthYear": 1881,
		},
		map[string]map[interface{}]interface{}{},
		0.5,
	},
	{
		rules2,
		map[string]interface{}{
			"city":      "Edirne",
			"birthYear": 1453,
		},
		map[string]map[interface{}]interface{}{},
		map[string]interface{}{
			"city":      "İstanbul",
			"birthYear": 1453,
		},
		map[string]map[interface{}]interface{}{},
		0,
	},
	{
		rules3,
		map[string]interface{}{
			"city":      "İstanbul",
			"birthYear": 1432,
			"skills":    []string{"tactics", "language", "technology", "art"},
		},
		map[string]map[interface{}]interface{}{},
		map[string]interface{}{
			"city":      "Selanik",
			"birthYear": 1881,
			"skills":    []string{"tactics", "language", "defence", "industry"},
		},
		map[string]map[interface{}]interface{}{},
		0.4444444444444444,
	},
	{
		rules4,
		map[string]interface{}{
			"city":      "İstanbul",
			"birthYear": 1432,
			"skills":    []string{"tactics", "language", "technology", "art"},
		},
		map[string]map[interface{}]interface{}{},
		map[string]interface{}{
			"city":      "Edirne",
			"birthYear": 1881,
			"skills":    []string{"tactics", "language", "defence", "industry"},
		},
		map[string]map[interface{}]interface{}{},
		0.5833333333333333,
	},
	{
		rules4,
		map[string]interface{}{
			"city":      "İstanbul",
			"birthYear": 1432,
			"skills":    []string{"tactics", "language", "technology", "art"},
		},
		map[string]map[interface{}]interface{}{},
		map[string]interface{}{
			"city":      "Selanik",
			"birthYear": 1881,
			"skills":    []string{"tactics", "language", "defence", "industry"},
		},
		map[string]map[interface{}]interface{}{},
		0.3333333333333333,
	},
	{
		rules4,
		map[string]interface{}{
			"city":      "İstanbul",
			"birthYear": 1432,
			"skills":    []string{"tactics", "language", "technology", "art"},
		},
		map[string]map[interface{}]interface{}{},
		map[string]interface{}{
			"city":      "Selanik",
			"birthYear": 1881,
			"skills":    []string{"tactics", "language", "defence", "industry"},
		},
		map[string]map[interface{}]interface{}{},
		0.3333333333333333,
	},
}

func TestCalculatePair(t *testing.T) {
	for i, test := range calculatePairTests {
		result := CalculatePair(test.rules, test.list1FieldOptionLabelMap, test.list2FieldOptionLabelMap, test.data1, test.data2)
		if result != test.expected {
			t.Errorf("Case %d: expected %v, got %v", i, test.expected, result)
		}
	}
}

// endregion
