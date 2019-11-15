package atama

import (
	"github.com/devingen/atama-api/model"
	"reflect"
	"testing"
)

type SeparateItemsByLimitsTest struct {
	items    []map[string]interface{}
	expected []MatchItem
}

var separateItemsByLimitsTests = []SeparateItemsByLimitsTest{
	{
		[]map[string]interface{}{
			{
				"email":     "fatih@sultan.com",
				"firstName": "Fatih",
				"lastName":  "Sultan",
				"birthYear": 1432,
			},
			{
				"email":     "ata@turk.com",
				"firstName": "Ata",
				"lastName":  "Türk",
				"birthYear": 1881,
				"city":      "Ankara",
				"limit":     2,
			},
		},
		[]MatchItem{
			{
				"_id":       "fatih@sultan.com:0",
				"_v":        0,
				"email":     "fatih@sultan.com",
				"firstName": "Fatih",
				"lastName":  "Sultan",
				"birthYear": 1432,
			},
			{
				"_id":       "ata@turk.com:0",
				"_v":        0,
				"email":     "ata@turk.com",
				"firstName": "Ata",
				"lastName":  "Türk",
				"birthYear": 1881,
				"city":      "Ankara",
				"limit":     2,
			},
			{
				"_id":       "ata@turk.com:1",
				"_v":        1,
				"email":     "ata@turk.com",
				"firstName": "Ata",
				"lastName":  "Türk",
				"birthYear": 1881,
				"city":      "Ankara",
				"limit":     2,
			},
		},
	},
}

var config = model.CalculatorConfig{
	FieldID:    "email",
	FieldLimit: "limit",
}

func TestSeparateItemsByLimits(t *testing.T) {
	for i, test := range separateItemsByLimitsTests {
		result := SeparateItemsByLimits(config, test.items)

		if len(result) != len(test.expected) {
			t.Errorf("Case %d: expected result to have %v items, got %v", i, len(test.expected), len(result))
		}

		for j, item := range result {
			expectedItem := test.expected[j]

			if !reflect.DeepEqual(item, expectedItem) {
				t.Errorf("Case %d", i)
				t.Errorf("Got     : %v", item)
				t.Errorf("Expected: %v", expectedItem)
			}
		}
	}
}
