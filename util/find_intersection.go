package util

func FindIntersection(array1, array2 []interface{}) []interface{} {

	array1Map := map[interface{}]bool{}
	for _, v := range array1 {
		array1Map[v] = true
	}

	intersection := make([]interface{}, 0)
	for _, v := range array2 {
		if array1Map[v] {
			intersection = append(intersection, v)
		}
	}

	return intersection
}
