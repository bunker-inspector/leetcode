package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type reverseIntegerCase struct {
	Input  int
	Answer int
}

func ReverseIntegerTest(t *testing.T) {
	for _, testCase := range []reverseIntegerCase{
		{
			Input:  123,
			Answer: 321,
		},
		{
			Input:  -123,
			Answer: -321,
		},
		{
			Input:  120,
			Answer: 21,
		},
		{
			Input:  0,
			Answer: 0,
		},
		{
			Input:  1534236469,
			Answer: 0,
		},
		{
			Input:  -2147483648,
			Answer: 0,
		},
	} {
		result := problems.ReverseInteger(testCase.Input)
		if result != testCase.Answer {
			t.Errorf("Expected %d, got %d", testCase.Answer, result)
			t.Fail()
		}
	}
}
