package atama

import (
	"github.com/devingen/atama-api/model"
	"sort"
)

type PairScore struct {
	Item  MatchItem `json:"item"`
	Score float64   `json:"score"`
}

type MatchItemScores struct {
	Item    MatchItem   `json:"item"`
	Matches []PairScore `json:"matches"`
}

func GenerateFieldOptionLabelMap(fields []model.GenericField) FieldOptionLabelMap {
	fieldOptionLabelMap := FieldOptionLabelMap{}
	for _, field := range fields {
		if field.HasOptions() {
			optionLabelMap := map[interface{}]interface{}{}
			for _, option := range field.GetOptions() {
				optionLabelMap[option.GetValue()] = option.GetLabel()
			}
			fieldOptionLabelMap[field.GetID()] = optionLabelMap
		}
	}
	return fieldOptionLabelMap
}

func GenerateScoreMatrix(
	rules []model.ConditionalComparisonRule,
	items1, items2 []MatchItem,
	list1Fields, list2Fields []model.GenericField,
) []MatchItemScores {

	list1FieldOptionLabelMap := GenerateFieldOptionLabelMap(list1Fields)
	list2FieldOptionLabelMap := GenerateFieldOptionLabelMap(list2Fields)

	matrix := make([]MatchItemScores, len(items1))
	for i, firstItem := range items1 {

		matches := make([]PairScore, 0)

		for _, secondItem := range items2 {
			score := CalculatePair(
				rules,
				list1FieldOptionLabelMap,
				list2FieldOptionLabelMap,
				firstItem,
				secondItem,
			)

			//if score != 0 {
			matches = append(matches, PairScore{
				Item:  secondItem,
				Score: score,
			})
			//}
		}

		sort.SliceStable(matches, func(i, j int) bool {
			return matches[i].Score > matches[j].Score
		})

		matrix[i] = MatchItemScores{
			Item:    firstItem,
			Matches: matches,
		}
	}

	return matrix
}
