package util

func ContainsItem(array1 []interface{}, item interface{}) bool {
	for _, value := range array1 {
		if value == item {
			return true
		}
	}
	return false
}
