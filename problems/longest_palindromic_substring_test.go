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
		longestPalindromicSubstringCase{
			Input:  "babad",
			Answer: "bab",
		},
		longestPalindromicSubstringCase{
			Input:  "cbbd",
			Answer: "bb",
		},
		longestPalindromicSubstringCase{
			Input:  "a",
			Answer: "a",
		},
		longestPalindromicSubstringCase{
			Input:  "ac",
			Answer: "a",
		},
		longestPalindromicSubstringCase{
			Input:  "acbcbca",
			Answer: "acbcbca",
		},
		longestPalindromicSubstringCase{
			Input:  "acbcbcacbc",
			Answer: "acbcbca",
		},
		longestPalindromicSubstringCase{
			Input:  "",
			Answer: "",
		},
		longestPalindromicSubstringCase{
			Input:  "vwxyz",
			Answer: "v",
		},
		longestPalindromicSubstringCase{
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
