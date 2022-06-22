package main

import (
	"reflect"
	"testing"
)

type Case struct {
	str   []string
	equal map[string][]string
}

func TestSort(t *testing.T) {
	testCases := []Case{
		{
			str:   []string{"пяТка", "ТЯПКА", "321", "22", "столик", "тест", "пятак"},
			equal: map[string][]string{"пятка": {"пятак", "тяпка"}},
		},
	}
	for _, v := range testCases {
		result := getAnagrams(v.str)
		if !reflect.DeepEqual(result, v.equal) {
			t.Errorf("For: %s. Expected: %s. Got: %s.", v.str, v.equal, result)
		}
	}
}
