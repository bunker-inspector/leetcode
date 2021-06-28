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
		reverseIntegerCase{
			Input:  123,
			Answer: 321,
		},
		reverseIntegerCase{
			Input:  -123,
			Answer: -321,
		},
		reverseIntegerCase{
			Input:  120,
			Answer: 21,
		},
		reverseIntegerCase{
			Input:  0,
			Answer: 0,
		},
		reverseIntegerCase{
			Input:  1534236469,
			Answer: 0,
		},
		reverseIntegerCase{
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
