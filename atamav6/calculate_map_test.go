package atamav6

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestCalculateMap(t *testing.T) {

	scoreMapMap := ScoreMapMap{
		101: {
			1: 25,
			2: 50,
		},
		102: {
			1: 25,
			2: 1,
		},
	}

	result := CalculateMap(len(scoreMapMap), scoreMapMap, PairMap{}, PairMap{}, 0)
	spew.Dump(result.PairMapSecondToFirst)
}
