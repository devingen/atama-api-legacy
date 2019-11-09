package util

import "testing"

type ContainsArrayTest struct {
	array1   []interface{}
	array2   []interface{}
	expected bool
}

var containsArrayTests = []ContainsArrayTest{
	{[]interface{}{1, 2}, []interface{}{2, 3}, false},
	{[]interface{}{1, 2}, []interface{}{3}, false},
	{[]interface{}{1, 2}, []interface{}{1}, true},
	{[]interface{}{1, 2}, []interface{}{2}, true},
	{[]interface{}{1, 2}, []interface{}{1, 2}, true},
}

func TestContainsArray(t *testing.T) {
	for i, test := range containsArrayTests {
		result := ContainsArray(test.array1, test.array2)
		if result != test.expected {
			t.Errorf("Case %d: expected %v, got %v", i, test.expected, result)
		}
	}
}
