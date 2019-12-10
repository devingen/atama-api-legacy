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
	"github.com/devingen/atama-api/atamav9"
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

	result := atamav1.CalculateMap(len(body.List1), scoreMapMap, atamav1.PairMap{}, atamav1.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 1       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, int(result.MaxScore*10))
}

func ExecuteTestV2(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav2.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav2.CalculateMap(len(body.List1), scoreMapMap, atamav2.PairMap{}, atamav2.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 2       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, int(result.MaxScore*10))
}

func ExecuteTestV3(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav3.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav3.CalculateMap(len(body.List1), scoreMapMap, atamav3.PairMap{}, atamav3.PairMap{}, 0)
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

	result := atamav5.CalculateMap(len(body.List1), scoreMapMap, atamav5.PairMap{}, atamav5.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 5       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, result.MaxScore)
}

func ExecuteTestV6(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav6.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav6.CalculateMap(len(body.List1), scoreMapMap, atamav6.PairMap{}, atamav6.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 6       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, result.MaxScore)
}

func ExecuteTestV7(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav7.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav7.CalculateMap(len(body.List1), scoreMapMap, atamav7.PairMap{}, atamav7.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 7       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, int(result.MaxScore*10))
}

func ExecuteTestV8(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav8.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav8.CalculateMap(len(body.List1), scoreMapMap, atamav8.PairMap{}, atamav8.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 8       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, result.MaxScore)
}

func ExecuteTestV9(body dto.BuildPairsBody) {
	start := time.Now()
	scoreMapMap, _ := atamav9.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	elapsedForGenerateScore := time.Since(start)

	result := atamav9.CalculateMap(len(body.List1), util.MaxIterationLimit(len(body.List1)), scoreMapMap, atamav9.PairMap{}, atamav9.PairMap{}, 0)
	elapsedForCalculatePair := time.Since(start)

	fmt.Printf("| 9       | %-13s | %-13s | %-18d | %-14d | %-5d |\n", elapsedForGenerateScore, elapsedForCalculatePair, result.PossibleMatchCount, result.IterationMatchCount, int(result.MaxScore*10))
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

		fmt.Println()
		fmt.Println("string to float64 map")
		for i < repetition {
			ExecuteTestV1(*body)
			i += 1
		}

		fmt.Println()
		fmt.Println("float64 to float64 map")
		i = 0
		for i < repetition {
			ExecuteTestV2(*body)
			i += 1
		}

		fmt.Println()
		fmt.Println("uint32 to uint32 map with IDs hashed into uint32")
		i = 0
		for i < repetition {
			ExecuteTestV3(*body)
			i += 1
		}

		fmt.Println()
		fmt.Println("uint32 to uint32 map with sequential uint32")
		i = 0
		for i < repetition {
			ExecuteTestV4(*body)
			i += 1
		}

		fmt.Println()
		fmt.Println("uint16 to uint16 map with sequential uint16")
		i = 0
		for i < repetition {
			ExecuteTestV5(*body)
			i += 1
		}

		fmt.Println()
		fmt.Println("int to int map with sequential int")
		i = 0
		for i < repetition {
			ExecuteTestV6(*body)
			i += 1
		}

		fmt.Println()
		fmt.Println("uint32 to uint32 map with IDs hashed into uint32 and scores stored as sorted stack")
		i = 0
		for i < repetition {
			ExecuteTestV7(*body)
			i += 1
		}

		fmt.Println()
		fmt.Println("int32 to int32 map with IDs hashed into int32 and scores stored as sorted stack")
		i = 0
		for i < repetition {
			ExecuteTestV8(*body)
			i += 1
		}

		fmt.Println()
		fmt.Println("uint32 to uint32 map with IDs hashed into uint32 and scores stored as sorted stack with limited iteration")
		i = 0
		for i < repetition {
			ExecuteTestV9(*body)
			i += 1
		}

		fmt.Println()
	}
}
