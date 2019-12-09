package dto

import (
	"github.com/devingen/atama-api/atama"
	"github.com/devingen/atama-api/model"
)

type BuildPairsBody struct {
	Rules       []model.ConditionalComparisonRule `json:"rules"`
	List1       []atama.MatchItem                 `json:"list1"`
	List1Fields []model.GenericField              `json:"list1Fields"`
	List2       []atama.MatchItem                 `json:"list2"`
	List2Fields []model.GenericField              `json:"list2Fields"`
}

type BuildPairsResponseBody struct {
	AverageScore         float64                `json:"averageScore"`
	PossibleMatchCount   int                    `json:"possibleMatchCount"`
	PairMapFirstToSecond model.PairScoreListMap `json:"pairMapFirstToSecond"`
	PairMapSecondToFirst model.PairScoreListMap `json:"pairMapSecondToFirst"`
	ScoreMap             model.ScoreMapMap      `json:"scoreMap"`
}
