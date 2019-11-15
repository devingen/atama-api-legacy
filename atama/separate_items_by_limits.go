package atama

import (
	"github.com/devingen/atama-api/model"
	"strconv"
)

func getItemID(config model.CalculatorConfig, item map[string]interface{}) string {
	id, hasID := item[config.FieldID].(string)
	if !hasID {
		return ""
	}
	return id
}

func getItemLimit(config model.CalculatorConfig, item map[string]interface{}) int {
	limit, hasLimit := item[config.FieldLimit].(int)
	if !hasLimit {
		return 1
	}
	return limit
}

type MatchItem map[string]interface{}

func (item MatchItem) GetID() string {
	return item["_id"].(string)
}

func (item MatchItem) GetVariant() int {
	return item["_variant"].(int)
}

func createMatchItem(config model.CalculatorConfig, item map[string]interface{}, variant int) MatchItem {
	matchItem := MatchItem{}
	for k, v := range item {
		matchItem[k] = v
	}
	matchItem["_id"] = getItemID(config, item) + ":" + strconv.Itoa(variant)
	matchItem["_v"] = variant
	return matchItem
}

func SeparateItemsByLimits(config model.CalculatorConfig, items []map[string]interface{}) []MatchItem {

	matchItems := make([]MatchItem, 0)
	for _, item := range items {
		limit := getItemLimit(config, item)
		for j := 0; j < limit; j++ {
			matchItems = append(matchItems, createMatchItem(config, item, j))
		}
	}
	return matchItems
}
