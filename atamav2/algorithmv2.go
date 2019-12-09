package atamav2

import (
	"github.com/devingen/atama-api/atama"
	"github.com/devingen/atama-api/model"
	"github.com/devingen/atama-api/util"
)

type ScoreMap map[float64]float64
type ScoreMapMap map[float64]ScoreMap
type HashIDMap map[string]float64
type PairMap map[float64]float64

func GenerateScoreMap(
	rules []model.ConditionalComparisonRule,
	items1, items2 []atama.MatchItem,
	list1Fields, list2Fields []model.GenericField,
) (ScoreMapMap, HashIDMap) {

	list1FieldOptionLabelMap := atama.GenerateFieldOptionLabelMap(list1Fields)
	list2FieldOptionLabelMap := atama.GenerateFieldOptionLabelMap(list2Fields)

	hashIDMap := HashIDMap{}
	scoreMapMap := ScoreMapMap{}
	for _, firstItem := range items1 {
		scoreMap := ScoreMap{}

		for _, secondItem := range items2 {
			score := atama.CalculatePair(
				rules,
				list1FieldOptionLabelMap,
				list2FieldOptionLabelMap,
				firstItem,
				secondItem,
			)

			if score > 0 {
				secondItemIdHash, hasHash := hashIDMap[secondItem.GetID()]
				if !hasHash {
					secondItemIdHash = float64(util.Hash(secondItem.GetID()))
					hashIDMap[secondItem.GetID()] = secondItemIdHash
				}
				scoreMap[secondItemIdHash] = score
			}
		}

		firstItemIdHash, hasHash := hashIDMap[firstItem.GetID()]
		if !hasHash {
			firstItemIdHash = float64(util.Hash(firstItem.GetID()))
			hashIDMap[firstItem.GetID()] = firstItemIdHash
		}
		scoreMapMap[firstItemIdHash] = scoreMap
	}
	return scoreMapMap, hashIDMap
}

func getFirstAvailableMatchInMap(pivotScoreMap ScoreMap, occupationMap PairMap) (float64, float64) {

	var maxScore float64
	var itemWithMaxScore float64
	for secondItemID, score := range pivotScoreMap {
		_, isOccupied := occupationMap[secondItemID]
		if !isOccupied && score > maxScore {
			maxScore = score
			itemWithMaxScore = secondItemID
		}
	}

	return itemWithMaxScore, maxScore
}

type CalculationResultMap struct {
	IterationMatchCount  int     `json:"iterationCount"`
	PossibleMatchCount   int     `json:"possibleMatchCount"`
	MaxScore             float64 `json:"maxScore"`
	PairMapFirstToSecond PairMap `json:"pairMapFirstToSecond"`
	PairMapSecondToFirst PairMap `json:"pairMapSecondToFirst"`
}

type SkipMap map[string]bool

func ClonePairMap(occupationMap PairMap) PairMap {
	clone := PairMap{}
	if occupationMap != nil {
		for k, v := range occupationMap {
			clone[k] = v
		}
	}
	return clone
}

func CalculateMap(maxLevel int, scoreMapMap ScoreMapMap, pairMapSecondToFirst PairMap, pairMapFirstToSecond PairMap, level int) CalculationResultMap {

	result := CalculationResultMap{
		IterationMatchCount:  1,
		PossibleMatchCount:   0,
		MaxScore:             0,
		PairMapSecondToFirst: pairMapSecondToFirst,
		PairMapFirstToSecond: pairMapFirstToSecond,
	}

	for firstItemID, scoreMap := range scoreMapMap {

		if _, hasID := pairMapFirstToSecond[firstItemID]; hasID {
			continue
		}

		internalPairMapSecondToFirst := ClonePairMap(pairMapSecondToFirst)
		internalPairMapFirstToSecond := ClonePairMap(pairMapFirstToSecond)

		secondItemId, score := getFirstAvailableMatchInMap(scoreMap, pairMapSecondToFirst)
		if secondItemId != 0 {
			result.PossibleMatchCount += 1
			internalPairMapFirstToSecond[firstItemID] = secondItemId
			internalPairMapSecondToFirst[secondItemId] = firstItemID
		}

		if level == maxLevel {
			break
		}

		innerResult := CalculateMap(maxLevel, scoreMapMap, internalPairMapSecondToFirst, internalPairMapFirstToSecond, level+1)

		result.IterationMatchCount += innerResult.IterationMatchCount
		result.PossibleMatchCount += innerResult.PossibleMatchCount
		score += innerResult.MaxScore

		if score >= result.MaxScore {
			result.MaxScore = score
			result.PairMapFirstToSecond = innerResult.PairMapFirstToSecond
			result.PairMapSecondToFirst = innerResult.PairMapSecondToFirst
		}
	}

	return result
}
