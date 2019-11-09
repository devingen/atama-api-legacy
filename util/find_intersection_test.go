package util

import "testing"

type FindIntersectionTest struct {
	array1   []interface{}
	array2   []interface{}
	expected []interface{}
}

var findIntersectionTests = []FindIntersectionTest{
	{[]interface{}{1, 2}, []interface{}{2, 3}, []interface{}{2}},
	{[]interface{}{1}, []interface{}{3}, []interface{}{}},
	{[]interface{}{1, 2, 3}, []interface{}{2, 3}, []interface{}{2, 3}},
	{[]interface{}{"a", "b", "c"}, []interface{}{"b", "c"}, []interface{}{"b", "c"}},
	{[]interface{}{"a", "b", "c"}, []interface{}{"b"}, []interface{}{"b"}},
	{[]interface{}{"a", "b"}, []interface{}{"c"}, []interface{}{}},
}

func TestFindIntersection(t *testing.T) {
	for i, test := range findIntersectionTests {
		result := FindIntersection(test.array1, test.array2)
		if !EqualsArray(result, test.expected) {
			t.Errorf("Case %d: expected %v, got %v", i, test.expected, result)
		}
	}
}
