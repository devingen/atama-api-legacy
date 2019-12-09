package atama

import (
	"strconv"
)

func getItemID(item map[string]interface{}) string {
	id, hasID := item["_id"].(string)
	if !hasID {
		return ""
	}
	return id
}

type MatchItem map[string]interface{}

func (item MatchItem) GetID() string {
	return item["_id"].(string)
}

func (item MatchItem) GetVariant() int {
	return item["_variant"].(int)
}

func (item MatchItem) GetLimit() int {
	limit, hasLimit := item["limit"].(int)
	if !hasLimit {
		return 1
	}
	return limit
}

func enrichMatchItem(item MatchItem, variant int) MatchItem {
	matchItem := MatchItem{}
	for k, v := range item {
		matchItem[k] = v
	}
	matchItem["_id"] = getItemID(item) + ":" + strconv.Itoa(variant)
	matchItem["_v"] = variant
	return matchItem
}

func SeparateItemsByLimits(itemMatchWrappers []MatchItemScores) []MatchItemScores {

	matchItems := make([]MatchItemScores, 0)
	for _, itemMatchWrapper := range itemMatchWrappers {
		limit := itemMatchWrapper.Item.GetLimit()
		for j := 0; j < limit; j++ {
			matchItems = append(matchItems, MatchItemScores{
				Item:    enrichMatchItem(itemMatchWrapper.Item, j),
				Matches: itemMatchWrapper.Matches,
			})
		}
	}
	return matchItems
}

func impoverishMatchItem(item MatchItem, variant int) MatchItem {
	matchItem := MatchItem{}
	for k, v := range item {
		matchItem[k] = v
	}
	matchItem["_id"] = getItemID(item) + ":" + strconv.Itoa(variant)
	matchItem["_v"] = variant
	return matchItem
}

func JoinSeparatedItems(separatedItemMatchWrappers []MatchItemScores) []MatchItemScores {

	matchItems := make([]MatchItemScores, 0)
	for _, itemMatchWrapper := range separatedItemMatchWrappers {
		limit := itemMatchWrapper.Item.GetLimit()
		for j := 0; j < limit; j++ {
			matchItems = append(matchItems, MatchItemScores{
				Item:    enrichMatchItem(itemMatchWrapper.Item, j),
				Matches: itemMatchWrapper.Matches,
			})
		}
	}
	return matchItems
}
