package util

func ConvertToNumber(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case float64:
		return int(v)
	}
	return 0
}
