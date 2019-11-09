package util

import (
	"fmt"
	"reflect"
)

func EqualsMap(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		fmt.Println(len(map1), len(map2))
		fmt.Println("a")
		return false
	}

	for k, v := range map1 {
		switch castedValue1 := v.(type) {
		case map[string]interface{}:
			switch castedValue2 := v.(type) {
			case map[string]interface{}:
				return EqualsMap(castedValue1, castedValue2)
			default:
				fmt.Println("b")
				return false
			}
		}

		if map2[k] != v {
			fmt.Println(map2[k], v)
			fmt.Println(reflect.TypeOf(map2[k]), reflect.TypeOf(v))
			fmt.Println("c")
			return false
		}
	}

	return true
}
