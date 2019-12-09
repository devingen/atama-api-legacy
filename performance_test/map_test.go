package performance

import (
	"fmt"
	"github.com/devingen/atama-api/atamav1"
	"github.com/devingen/atama-api/atamav2"
	"github.com/devingen/atama-api/atamav3"
	"github.com/devingen/atama-api/atamav4"
	"github.com/devingen/atama-api/atamav5"
	"github.com/devingen/atama-api/atamav6"
	"github.com/devingen/atama-api/atamav7"
	"github.com/devingen/atama-api/atamav8"
	"github.com/devingen/atama-api/dto"
	"github.com/devingen/atama-api/util"
	"testing"
	"time"
)

const repetition = 5

func ExecuteTestV1(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav1.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav1.CalculateMap(10, scoreMapMap, atamav1.PairMap{}, atamav1.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 1       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, int(result.MaxScore*10))
}

func ExecuteTestV2(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav2.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav2.CalculateMap(10, scoreMapMap, atamav2.PairMap{}, atamav2.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 2       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, int(result.MaxScore*10))
}

func ExecuteTestV3(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav3.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav3.CalculateMap(10, scoreMapMap, atamav3.PairMap{}, atamav3.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 3       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, result.MaxScore)
}

func ExecuteTestV4(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav4.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav4.CalculateMap(10, scoreMapMap, atamav4.PairMap{}, atamav4.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 4       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, result.MaxScore)
}

func ExecuteTestV5(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav5.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav5.CalculateMap(10, scoreMapMap, atamav5.PairMap{}, atamav5.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 5       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, result.MaxScore)
}

func ExecuteTestV6(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav6.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav6.CalculateMap(10, scoreMapMap, atamav6.PairMap{}, atamav6.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 6       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, result.MaxScore)
}

func ExecuteTestV7(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav7.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav7.CalculateMap(len(body.List1), util.MaxIterationLimit(len(body.List1)), scoreMapMap, atamav7.PairMap{}, atamav7.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 7       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, int(result.MaxScore*10))
}

func ExecuteTestV8(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav8.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav8.CalculateMap(10, scoreMapMap, atamav8.PairMap{}, atamav8.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 8       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, result.MaxScore)
}

func TestPerformanceWithMap(t *testing.T) {

	for _, file := range files {

		body, err := ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		m := len(body.List1)
		n := len(body.List2)

		fmt.Printf("File:                  %s\n", file)
		fmt.Printf("First item count:      %d\n", m)
		fmt.Printf("Second item count:     %d\n", n)
		fmt.Println()
		fmt.Println("| Version | GenerateScore | CalculatePair | PossibleMatchCount | IterationCount | Score |")
		fmt.Println("|---------|---------------|---------------|--------------------|----------------|-------|")

		i := 0

		//for i < repetition {
		//	ExecuteTestV1(*body)
		//	i += 1
		//}
		//
		//fmt.Println()
		//i = 0
		//for i < repetition {
		//	ExecuteTestV2(*body)
		//	i += 1
		//}

		//fmt.Println()
		//i = 0
		//for i < repetition {
		//	ExecuteTestV3(*body)
		//	i += 1
		//}

		//fmt.Println()
		//i = 0
		//for i < repetition {
		//	ExecuteTestV4(*body)
		//	i += 1
		//}
		//
		//fmt.Println()
		//i = 0
		//for i < repetition {
		//	ExecuteTestV5(*body)
		//	i += 1
		//}
		//
		//fmt.Println()
		//i = 0
		//for i < repetition {
		//	ExecuteTestV6(*body)
		//	i += 1
		//}

		fmt.Println()
		i = 0
		for i < repetition {
			ExecuteTestV7(*body)
			i += 1
		}

		//fmt.Println()
		//i = 0
		//for i < repetition {
		//	ExecuteTestV8(*body)
		//	i += 1
		//}

		fmt.Println()
	}
}
