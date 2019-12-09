package atamav1

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestCalculateMap(t *testing.T) {

	scoreMapMap := ScoreMapMap{
		"A": {
			"a": 25,
			"b": 50,
		},
		"B": {
			"a": 25,
			"b": 1,
		},
	}

	result := CalculateMap(len(scoreMapMap), scoreMapMap, PairMap{}, PairMap{}, 0)
	spew.Dump(result.PairMapSecondToFirst)
}
