package performance

import (
	"fmt"
	"github.com/devingen/atama-api/atama"
	"github.com/devingen/atama-api/util"
	"testing"
	"time"
)

func TestPerformanceWithList(t *testing.T) {

	for _, file := range files {

		body, err := ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		m := len(body.List1)
		n := len(body.List2)

		maxIterationLimit := util.MaxIterationLimit(m)
		maxIterationLevel := util.MaxIterationLevel(n)
		fmt.Println(maxIterationLimit, maxIterationLevel)

		start := time.Now()
		fmt.Printf("CalculateList %s", file)

		scoreMatrix := atama.GenerateScoreMatrix(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
		fmt.Printf(" %s", time.Since(start))

		result := atama.CalculateList(scoreMatrix, nil, maxIterationLimit, maxIterationLevel, 0)
		fmt.Printf(" %s", time.Since(start))
		fmt.Println()
		fmt.Println("CalculateList PMS: ", result.PossibleMatchCount)
		fmt.Println("CalculateList Score:  ", result.MaxScore)
	}
}
