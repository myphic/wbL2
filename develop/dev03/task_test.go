package main

import "testing"

type Case struct {
	str      []string
	equal    []string
	function func([]string) []string
}

func SliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	testCases := []Case{
		{
			str:      []string{"-1000", "2", "1", "1", "3", "4", "5", "132"},
			equal:    []string{"-1000", "1", "1", "2", "3", "4", "5", "132"},
			function: sortNumeric,
		},
		{
			str:      []string{"a", "b", "c"},
			equal:    []string{"c", "b", "a"},
			function: sortReverse,
		},
		{
			str:      []string{"a", "b", "b", "c"},
			equal:    []string{"a", "b", "c"},
			function: sortWithoutDuplicates,
		},
	}
	for _, v := range testCases {
		result := v.function(v.str)
		if !SliceEqual(result, v.equal) {
			t.Errorf("For: %s. Expected: %s. Got: %s.", v.str, v.equal, result)
		}
	}
}
