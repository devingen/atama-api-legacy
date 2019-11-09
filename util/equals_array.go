package util

func EqualsArray(array1, array2 []interface{}) bool {
	if len(array1) != len(array2) {
		return false
	}

	intersection := FindIntersection(array1, array2)
	if len(intersection) != len(array2) {
		return false
	}

	return true
}
