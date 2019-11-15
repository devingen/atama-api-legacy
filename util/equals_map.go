package util

func EqualsMap(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}

	for k, v := range map1 {
		switch castedValue1 := v.(type) {
		case map[string]interface{}:
			switch castedValue2 := v.(type) {
			case map[string]interface{}:
				return EqualsMap(castedValue1, castedValue2)
			default:
				return false
			}
		}

		if map2[k] != v {
			return false
		}
	}

	return true
}
