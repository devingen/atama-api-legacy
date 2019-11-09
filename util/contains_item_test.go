package util

import "testing"

type ContainsItemTest struct {
	array    []interface{}
	item     interface{}
	expected bool
}

var containsItemTests = []ContainsItemTest{
	{[]interface{}{1, 2}, 3, false},
	{[]interface{}{1, 2}, 1, true},
	{[]interface{}{1, 2}, 2, true},
	{[]interface{}{"a", "b"}, "c", false},
	{[]interface{}{"a", "b"}, "a", true},
	{[]interface{}{"a", "b"}, "b", true},
}

func TestContainsItem(t *testing.T) {
	for i, test := range containsItemTests {
		result := ContainsItem(test.array, test.item)
		if result != test.expected {
			t.Errorf("Case %d: expected %v, got %v", i, test.expected, result)
		}
	}
}
