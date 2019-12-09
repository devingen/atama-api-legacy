package atama

type OccupationMap map[string]string

func splice(slice []MatchItemScores, s int) []MatchItemScores {
	return append(slice[:s], slice[s+1:]...)
}

func getFirstAvailableMatch(pivot MatchItemScores, occupationMap OccupationMap) *PairScore {

	for _, match := range pivot.Matches {
		candidate := match.Item
		candidatesExistingMatch := occupationMap[candidate.GetID()]

		if candidatesExistingMatch == "" {
			// TODO multiple match between these items with their other variations
			return &match
		}
	}

	return nil
}

type ResultPairScore struct {
	ItemID string  `json:"itemId"`
	Score  float64 `json:"score"`
}

type CalculationResult struct {
	PossibleMatchCount int                        `json:"possibleMatchCount"`
	MaxScore           float64                    `json:"maxScore"`
	Matches            map[string]ResultPairScore `json:"matches"`
}

// possibleMatchCount = (maxIterationLimit + 1) ^ maxIterationLevel
// EXCEPTION maxIterationLevel=0 -> possibleMatchCount=1
//var maxIterationLimit = 6
//var maxIterationLevel = 6

// 8 should be 8. sqrt(8)=3
// 8 should be 8. sqrt(8/3)=1.6

// 40 should be 6, sqrt(40)=6
// 130 should be 6, sqrt(130/3)=6

func CalculateList(scoreMatrix []MatchItemScores, occupationMap OccupationMap, maxIterationLimit, maxIterationLevel int, level int) CalculationResult {

	var maxScore float64 = 0
	totalPossibleMatchCount := 0
	bestMatches := map[string]ResultPairScore{}

	//partLength := len(scoreMatrix) / maxIterationLimit
	for i := range scoreMatrix {

		index := i // * partLength
		//if index >= len(scoreMatrix) {
		//	index -= 1
		//}

		matches := map[string]ResultPairScore{}

		internalOccupation := map[string]string{}
		if occupationMap != nil {
			for k, v := range occupationMap {
				internalOccupation[k] = v
			}
		}

		scoreMatrixClone := make([]MatchItemScores, len(scoreMatrix))
		copy(scoreMatrixClone[:], scoreMatrix)

		pivot := scoreMatrixClone[index]

		var score float64 = 0
		match := getFirstAvailableMatch(pivot, internalOccupation)

		if match != nil {

			score = match.Score

			// add pivot's match to occupation map
			internalOccupation[match.Item.GetID()] = pivot.Item.GetID()

			// add pivot's match to match list
			matches[pivot.Item.GetID()] = ResultPairScore{
				ItemID: match.Item.GetID(),
				Score:  match.Score,
			}
		}

		if len(scoreMatrixClone) == 1 {
			// last node reached, skip the rest
			bestMatches = matches
			totalPossibleMatchCount += 1
			continue
		}

		// remove the pivot from the list and calculate the list for rest of the items
		rest := splice(scoreMatrixClone, index)

		//fmt.Println()
		//fmt.Println("level", level)
		//fmt.Println("pivot", pivot.Item.GetID())
		//for _, item := range rest {
		//	fmt.Println("      ", item.Item.GetID())
		//}

		innerCalculation := CalculateList(rest, internalOccupation, maxIterationLimit, maxIterationLevel, level+1)

		// add matches from the inner calculation
		for k, v := range innerCalculation.Matches {
			matches[k] = v
		}

		totalPossibleMatchCount += innerCalculation.PossibleMatchCount
		score += innerCalculation.MaxScore

		//if level == 0 {
		//	fmt.Println(index, "score", score)
		//}

		if score >= maxScore {
			bestMatches = matches
			maxScore = score
		}

		//if i >= maxIterationLimit || level >= maxIterationLevel {
		//	break
		//}
	}

	//if level == 0 {
	//	fmt.Println("totalPossibleMatchCount", totalPossibleMatchCount)
	//}

	return CalculationResult{
		PossibleMatchCount: totalPossibleMatchCount,
		MaxScore:           maxScore,
		Matches:            bestMatches,
	}
}
