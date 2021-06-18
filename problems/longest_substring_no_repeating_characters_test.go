package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type LongestNonrepeatingSubstringInput struct {
	S      string
	Answer int
}

func TestLongestNonrepeatingSubstring(t *testing.T) {
	for _, testCase := range []LongestNonrepeatingSubstringInput{
		LongestNonrepeatingSubstringInput{
			S:      "abcabcbb",
			Answer: 3,
		},
		LongestNonrepeatingSubstringInput{
			S:      "bbbbb",
			Answer: 1,
		},
		LongestNonrepeatingSubstringInput{
			S:      "pwwkew",
			Answer: 3,
		},
		LongestNonrepeatingSubstringInput{
			S:      "",
			Answer: 0,
		},
		LongestNonrepeatingSubstringInput{
			S:      "aab",
			Answer: 2,
		},
		LongestNonrepeatingSubstringInput{
			S:      "abba",
			Answer: 2,
		},
	} {
		result := problems.LongestNonrepeatingSubstring(testCase.S)
		if result != testCase.Answer {
			t.Errorf("Given input %s:\n Expected %d, got %d.\n", testCase.S, testCase.Answer, result)
		}
	}
}
