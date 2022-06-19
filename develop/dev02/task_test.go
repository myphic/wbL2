package main

import "testing"

type Case struct {
	value string
	equal string
}

func TestUnpackString(t *testing.T) {
	testCases := []Case{
		{
			value: "45",
			equal: "",
		},
		{
			value: "a4bc2d5e",
			equal: "aaaabccddddde",
		},
		{
			value: "abcd",
			equal: "abcd",
		},
		{
			value: "",
			equal: "",
		},
		{
			value: "abcd5tyuioop",
			equal: "abcdddddtyuioop",
		},
	}
	for _, v := range testCases {
		result, _ := UnpackString(v.value)
		if result != v.equal {
			t.Errorf("For: %s. Expected: %s. Got: %s.", v.value, v.equal, result)
		}
	}
}
