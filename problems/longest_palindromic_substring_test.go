package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type longestPalindromicSubstringCase struct {
	Input  string
	Answer string
}

func TestLongestPalindrome(t *testing.T) {
	for _, testCase := range []longestPalindromicSubstringCase{
		{
			Input:  "babad",
			Answer: "bab",
		},
		{
			Input:  "cbbd",
			Answer: "bb",
		},
		{
			Input:  "a",
			Answer: "a",
		},
		{
			Input:  "ac",
			Answer: "a",
		},
		{
			Input:  "acbcbca",
			Answer: "acbcbca",
		},
		{
			Input:  "acbcbcacbc",
			Answer: "acbcbca",
		},
		{
			Input:  "",
			Answer: "",
		},
		{
			Input:  "vwxyz",
			Answer: "v",
		},
		{
			Input:  "ddacbcbca",
			Answer: "acbcbca",
		},
	} {
		result := problems.LongestPalindrome(testCase.Input)
		if result != testCase.Answer {
			t.Errorf("Given input %s: \n Expected %s, got %s.\n",
				testCase.Input,
				testCase.Answer,
				result)
		}
	}
}
