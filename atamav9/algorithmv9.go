package atamav9

import (
	"github.com/devingen/atama-api/atama"
	"github.com/devingen/atama-api/model"
	"github.com/devingen/atama-api/util"
	"sort"
)

type ScoreStack []uint32
type ScoreMap map[uint32]float64

type ScoreBundle struct {
	ScoreStack ScoreStack
	ScoreMap   ScoreMap
}

type ScoreStackMap map[uint32]ScoreBundle
type PairMap map[uint32]uint32

type IDHashMap map[string]uint32
type HashIDMap map[uint32]string

func GenerateScoreMap(
	rules []model.ConditionalComparisonRule,
	items1, items2 []atama.MatchItem,
	list1Fields, list2Fields []model.GenericField,
) (ScoreStackMap, HashIDMap) {

	list1FieldOptionLabelMap := atama.GenerateFieldOptionLabelMap(list1Fields)
	list2FieldOptionLabelMap := atama.GenerateFieldOptionLabelMap(list2Fields)

	IDHashMap := IDHashMap{}
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
				secondItemIdHash, hasHash := IDHashMap[secondItem.GetID()]
				if !hasHash {
					secondItemIdHash = util.Hash(secondItem.GetID())
					IDHashMap[secondItem.GetID()] = secondItemIdHash
					hashIDMap[secondItemIdHash] = secondItem.GetID()
				}
				scoreMap[secondItemIdHash] = score
				scoreStack = append(scoreStack, secondItemIdHash)
			}
		}

		firstItemIdHash, hasHash := IDHashMap[firstItem.GetID()]
		if !hasHash {
			firstItemIdHash = util.Hash(firstItem.GetID())
			IDHashMap[firstItem.GetID()] = firstItemIdHash
			hashIDMap[firstItemIdHash] = firstItem.GetID()
		}
		sort.Slice(scoreStack, func(a, b int) bool {
			return scoreMap[uint32(a)] > scoreMap[uint32(b)]
		})
		scoreMapMap[firstItemIdHash] = ScoreBundle{
			ScoreStack: scoreStack,
			ScoreMap:   scoreMap,
		}
	}
	return scoreMapMap, hashIDMap
}

func getFirstAvailableMatchInMap(scoreBundle ScoreBundle, occupationMap PairMap) (uint32, float64) {

	var maxScore float64
	var itemWithMaxScore uint32
	for _, secondItemID := range scoreBundle.ScoreStack {
		_, isOccupied := occupationMap[secondItemID]
		if !isOccupied {
			maxScore = scoreBundle.ScoreMap[secondItemID]
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

func CalculateMap(maxLevel, iterationLimit int, scoreMapMap ScoreStackMap, pairMapSecondToFirst PairMap, pairMapFirstToSecond PairMap, level int) CalculationResultMap {

	result := CalculationResultMap{
		IterationMatchCount:  1,
		PossibleMatchCount:   0,
		MaxScore:             0,
		PairMapSecondToFirst: pairMapSecondToFirst,
		PairMapFirstToSecond: pairMapFirstToSecond,
	}

	iteration := 0
	for firstItemID, scoreBundle := range scoreMapMap {

		if iteration > iterationLimit {
			break
		}
		iteration++

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

		innerResult := CalculateMap(maxLevel, iterationLimit, scoreMapMap, internalPairMapSecondToFirst, internalPairMapFirstToSecond, level+1)

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
