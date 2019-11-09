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
	ItemID string
	Score  float64
}

type CalculationResult struct {
	possibleMatchCount int
	maxScore           float64
	matches            map[string]ResultPairScore
}

func CalculateList(scoreMatrix []MatchItemScores, occupationMap OccupationMap) CalculationResult {

	if occupationMap == nil {
		occupationMap = OccupationMap{}
	}

	var maxScore float64 = 0
	totalPossibleMatchCount := 0
	matches := map[string]ResultPairScore{}

	for i := range scoreMatrix {
		pivot := scoreMatrix[i]
		match := getFirstAvailableMatch(pivot, occupationMap)

		if match == nil {

			if len(scoreMatrix) == 1 {
				return CalculationResult{
					possibleMatchCount: 0,
					maxScore:           0,
					matches:            matches,
				}
			}
			continue
		}

		rest := splice(scoreMatrix, i)

		occupationMap[match.Item.GetID()] = pivot.Item.GetID()

		innerCalculation := CalculateList(rest, occupationMap)

		totalPossibleMatchCount += innerCalculation.possibleMatchCount
		totalScore := innerCalculation.maxScore + match.Score

		if totalScore > maxScore {
			maxScore = totalScore

			for k, v := range innerCalculation.matches {
				matches[k] = v
			}
			matches[pivot.Item.GetID()] = ResultPairScore{
				ItemID: match.Item.GetID(),
				Score:  match.Score,
			}
			totalPossibleMatchCount++
		}
	}

	return CalculationResult{
		possibleMatchCount: totalPossibleMatchCount,
		maxScore:           maxScore,
		matches:            matches,
	}
}
