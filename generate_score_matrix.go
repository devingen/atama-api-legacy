package atama

import (
	"github.com/devingen/atama-api/model"
	"sort"
)

type PairScore struct {
	Item  MatchItem
	Score float64
}

type MatchItemScores struct {
	Item    MatchItem
	Matches []PairScore
}

func GenerateScoreMatrix(config model.CalculatorConfig, rules []model.ConditionalComparisonRule, items1, items2 []MatchItem) []MatchItemScores {

	matrix := make([]MatchItemScores, len(items1))
	for i, firstItem := range items1 {

		matches := make([]PairScore, len(items2))

		for j, secondItem := range items2 {
			matches[j] = PairScore{
				Item:  secondItem,
				Score: CalculatePair(config, rules, firstItem, secondItem),
			}
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
