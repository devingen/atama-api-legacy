package atama

import (
	"fmt"
	"testing"
)

type CalculateListTest struct {
	matrix []MatchItemScores
}

var calculateListTests = []CalculateListTest{
	{
		matrix: []MatchItemScores{
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

func TestCalculateList(t *testing.T) {
	for _, test := range calculateListTests {
		result := CalculateList(test.matrix, nil)

		fmt.Println(result)
		fmt.Println(result.possibleMatchCount)
		fmt.Println(len(result.matches))
		fmt.Println(result.maxScore)
		for pivotId, match := range result.matches {
			fmt.Println(pivotId, match.ItemID, match.Score)
		}
	}
}
