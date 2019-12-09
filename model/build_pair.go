package model

// used to send automatically matched pairs with scores
type PairScore struct {
	ID    string  `json:"id"`
	Score float64 `json:"score"`
}
type PairScoreListMap map[string][]PairScore

// used to send all the calculated scores in response
type ScoreMap map[string]float64
type ScoreMapMap map[string]ScoreMap
