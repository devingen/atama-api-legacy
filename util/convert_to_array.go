package util

func ConvertToInterfaceArray(value interface{}) []interface{} {
	switch v := value.(type) {
	case []interface{}:
		return v
	case []string:
		array := make([]interface{}, len(v))
		for i, itemValue := range v {
			array[i] = itemValue
		}
		return array
	case []float64:
		array := make([]interface{}, len(v))
		for i, itemValue := range v {
			array[i] = itemValue
		}
		return array
	case []int:
		array := make([]interface{}, len(v))
		for i, itemValue := range v {
			array[i] = itemValue
		}
		return array
	case []map[string]interface{}:
		array := make([]interface{}, len(v))
		for i, itemValue := range v {
			array[i] = itemValue
		}
		return array
	}
	return make([]interface{}, 0)
}
