package main

import (
	"testing"
)

func Test_fatRateSuggestion_GetSuggestion(t *testing.T) {
	sugg := GetFatRateSuggestion()
	tests := []Person{
		{
			sex:     "男",
			age:     35,
			fatRate: 0.24,
		},
	}
	if got := sugg.GetSuggestion(&tests[0]); got != "肥胖" {
		t.Fail()
	}
}
