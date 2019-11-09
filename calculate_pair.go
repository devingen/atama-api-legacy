package atama

import (
	"github.com/devingen/atama-api/model"
	"github.com/devingen/atama-api/util"
)

func calculateSimilarity(array1, array2 []interface{}) float64 {
	intersection := util.FindIntersection(array1, array2)
	total := len(array1) + len(array2) - len(intersection)
	return float64(len(intersection)) / float64(total)
}

func compareValues(comparison model.Comparison, value1, value2 interface{}) float64 {

	switch comparison {
	case model.ComparisonEq:
		if value1 == value2 {
			return 1
		}
	case model.ComparisonNe:
		if value1 != value2 {
			return 1
		}
	case model.ComparisonLt:
		if util.ConvertToNumber(value1) < util.ConvertToNumber(value2) {
			return 1
		}
	case model.ComparisonLte:
		if util.ConvertToNumber(value1) <= util.ConvertToNumber(value2) {
			return 1
		}
	case model.ComparisonGt:
		if util.ConvertToNumber(value1) > util.ConvertToNumber(value2) {
			return 1
		}
	case model.ComparisonGte:
		if util.ConvertToNumber(value1) >= util.ConvertToNumber(value2) {
			return 1
		}
	case model.ComparisonIn:
		if util.ContainsItem(util.ConvertToInterfaceArray(value2), value1) {
			return 1
		}
	case model.ComparisonSimilar:
		return calculateSimilarity(util.ConvertToInterfaceArray(value1), util.ConvertToInterfaceArray(value2))
	case model.ComparisonDifferent:
		return 1 - calculateSimilarity(util.ConvertToInterfaceArray(value1), util.ConvertToInterfaceArray(value2))
	case model.ComparisonContain:
		if util.ContainsArray(util.ConvertToInterfaceArray(value1), util.ConvertToInterfaceArray(value2)) {
			return 1
		}
	case model.ComparisonNcontain:
		if !util.ContainsArray(util.ConvertToInterfaceArray(value1), util.ConvertToInterfaceArray(value2)) {
			return 1
		}
	}
	return 0
}

func CalculatePair(config model.CalculatorConfig, rules []model.ConditionalComparisonRule, data1 map[string]interface{}, data2 map[string]interface{}) float64 {

	var points float64 = 0
	var rulesThatApply = 0

	for _, rule := range rules {

		if rule.Type == model.RuleTypeComparison {

			ruleDetails := rule.Comparisons[0]
			if ruleDetails.FirstField == nil || ruleDetails.Comparison == "" || ruleDetails.SecondField == nil {
				rulesThatApply += 1
				continue
			}

			//dataHolder1 := data1[config.FirstDataGroupIndex].(map[string]interface{})
			value1 := data1[ruleDetails.FirstField.ID]

			//dataHolder2 := data2[config.FirstDataGroupIndex].(map[string]interface{})
			value2 := data2[ruleDetails.SecondField.ID]

			comparisonPoints := compareValues(ruleDetails.Comparison, value1, value2)
			points = points + comparisonPoints;
			rulesThatApply += 1

		} else if rule.Type == model.RuleTypeConditionalComparison {

			// check if first comparison data is valid
			if rule.Comparisons[0].FirstField == nil || rule.Comparisons[0].Comparison == "" || rule.Comparisons[0].SecondValue == nil {
				rulesThatApply += 1
				continue
			}

			// check if second comparison data is valid
			if rule.Comparisons[1].FirstField == nil || rule.Comparisons[1].Comparison == "" || rule.Comparisons[1].SecondValue == nil {
				rulesThatApply += 1
				continue
			}

			//dataHolder1 := data1[config.FirstDataGroupIndex].(map[string]interface{})
			value1 := data1[rule.Comparisons[0].FirstField.ID]

			doesFirstConditionSatisfy := compareValues(
				rule.Comparisons[0].Comparison,
				value1,
				rule.Comparisons[0].SecondValue,
			) != 0

			if !doesFirstConditionSatisfy {
				// no need to check the second comparison if the first comparison doesn't pass
				continue
			}

			//dataHolder2 := data2[config.FirstDataGroupIndex].(map[string]interface{})
			value2 := data2[rule.Comparisons[1].FirstField.ID]

			comparisonPoints := compareValues(
				rule.Comparisons[1].Comparison,
				value2,
				rule.Comparisons[1].SecondValue,
			)

			points += comparisonPoints
			rulesThatApply += 1
		}
	}

	if points == 0 {
		return 0
	}
	return float64(points) / float64(len(rules))
}
