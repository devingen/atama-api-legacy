package atamav8

import (
	"github.com/devingen/atama-api/atama"
	"github.com/devingen/atama-api/model"
	"github.com/devingen/atama-api/util"
	"sort"
)

type ScoreStack []int32
type ScoreMap map[int32]uint16

type ScoreBundle struct {
	scoreStack ScoreStack
	scoreMap   ScoreMap
}

type ScoreStackMap map[int32]ScoreBundle
type PairMap map[int32]int32

type HashIDMap map[string]int32

func GenerateScoreMap(
	rules []model.ConditionalComparisonRule,
	items1, items2 []atama.MatchItem,
	list1Fields, list2Fields []model.GenericField,
) (ScoreStackMap, HashIDMap) {

	list1FieldOptionLabelMap := atama.GenerateFieldOptionLabelMap(list1Fields)
	list2FieldOptionLabelMap := atama.GenerateFieldOptionLabelMap(list2Fields)

	hashIDMap := HashIDMap{}
	scoreMapMap := ScoreStackMap{}
	for _, firstItem := range items1 {
		scoreMap := ScoreMap{}
		scoreStack := ScoreStack{}

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
					secondItemIdHash = int32(util.Hash(secondItem.GetID()))
					hashIDMap[secondItem.GetID()] = secondItemIdHash
				}
				scoreMap[secondItemIdHash] = uint16(score * 10)
				scoreStack = append(scoreStack, secondItemIdHash)
			}
		}

		firstItemIdHash, hasHash := hashIDMap[firstItem.GetID()]
		if !hasHash {
			firstItemIdHash = int32(util.Hash(firstItem.GetID()))
			hashIDMap[firstItem.GetID()] = firstItemIdHash
		}
		sort.Slice(scoreStack, func(a, b int) bool {
			return scoreMap[int32(a)] > scoreMap[int32(b)]
		})
		scoreMapMap[firstItemIdHash] = ScoreBundle{
			scoreStack: scoreStack,
			scoreMap:   scoreMap,
		}
	}
	return scoreMapMap, hashIDMap
}

func getFirstAvailableMatchInMap(scoreBundle ScoreBundle, occupationMap PairMap) (int32, uint16) {

	var maxScore uint16
	var itemWithMaxScore int32
	for _, secondItemID := range scoreBundle.scoreStack {
		_, isOccupied := occupationMap[secondItemID]
		if !isOccupied {
			maxScore = scoreBundle.scoreMap[secondItemID]
			itemWithMaxScore = secondItemID
		}
	}

	return itemWithMaxScore, maxScore
}

type CalculationResultMap struct {
	IterationMatchCount  int     `json:"iterationCount"`
	PossibleMatchCount   int     `json:"possibleMatchCount"`
	MaxScore             uint16  `json:"maxScore"`
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

func CalculateMap(maxLevel int, scoreMapMap ScoreStackMap, pairMapSecondToFirst PairMap, pairMapFirstToSecond PairMap, level int) CalculationResultMap {

	result := CalculationResultMap{
		IterationMatchCount:  1,
		PossibleMatchCount:   0,
		MaxScore:             0,
		PairMapSecondToFirst: pairMapSecondToFirst,
		PairMapFirstToSecond: pairMapFirstToSecond,
	}

	for firstItemID, scoreBundle := range scoreMapMap {

		if _, hasID := pairMapFirstToSecond[firstItemID]; hasID {
			continue
		}

		internalPairMapSecondToFirst := ClonePairMap(pairMapSecondToFirst)
		internalPairMapFirstToSecond := ClonePairMap(pairMapFirstToSecond)

		secondItemId, score := getFirstAvailableMatchInMap(scoreBundle, pairMapSecondToFirst)
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
