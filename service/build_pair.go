package service

import (
	"github.com/devingen/atama-api/atamav7"
	"github.com/devingen/atama-api/dto"
	"github.com/devingen/atama-api/model"
	"github.com/devingen/atama-api/util"
	"log"
	"time"
)

func BuildPairs(body dto.BuildPairsBody) dto.BuildPairsResponseBody {
	m := len(body.List1)
	n := len(body.List2)

	start := time.Now()
	log.Printf("")
	log.Printf("%d %d", m, n)

	scoreStackMap, hashIdMap := atamav7.GenerateScoreMap(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	log.Printf("GenerateScoreMap took %s", time.Since(start))

	result := atamav7.CalculateMap(len(body.List1), util.MaxIterationLimit(len(body.List1)), scoreStackMap, atamav7.PairMap{}, atamav7.PairMap{}, 0)
	log.Printf("CalculateMap took %s", time.Since(start))
	//spew.Dump(result)

	response := dto.BuildPairsResponseBody{
		AverageScore:         result.MaxScore,
		PossibleMatchCount:   result.PossibleMatchCount,
		PairMapFirstToSecond: model.PairScoreListMap{},
		PairMapSecondToFirst: model.PairScoreListMap{},
		ScoreMap:             model.ScoreMapMap{},
	}

	// build PairMapFirstToSecond with real IDs
	for firstUserIdHash, secondUserIdHash := range result.PairMapFirstToSecond {
		firstUserId := hashIdMap[firstUserIdHash]
		secondUserId := hashIdMap[secondUserIdHash]
		response.PairMapFirstToSecond[firstUserId] = []model.PairScore{{
			ID:    secondUserId,
			Score: scoreStackMap[firstUserIdHash].ScoreMap[secondUserIdHash],
		}}
	}

	// build PairMapFirstToSecond with real IDs
	for secondUserIdHash, firstUserIdHash := range result.PairMapSecondToFirst {
		firstUserId := hashIdMap[firstUserIdHash]
		secondUserId := hashIdMap[secondUserIdHash]
		response.PairMapSecondToFirst[secondUserId] = []model.PairScore{{
			ID:    firstUserId,
			Score: scoreStackMap[firstUserIdHash].ScoreMap[secondUserIdHash],
		}}
	}

	// build ScoreMapMap with real IDs
	for firstUserIdHash, scoreBundle := range scoreStackMap {
		firstUserId := hashIdMap[firstUserIdHash]
		response.ScoreMap[firstUserId] = model.ScoreMap{}
		for secondUserIdHash, score := range scoreBundle.ScoreMap {
			secondUserId := hashIdMap[secondUserIdHash]
			response.ScoreMap[firstUserId][secondUserId] = score
		}
	}

	return response
}
