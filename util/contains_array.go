package util

func ContainsArray(array1, array2 []interface{}) bool {
	intersection := FindIntersection(array1, array2)
	return len(intersection) == len(array2)
}
