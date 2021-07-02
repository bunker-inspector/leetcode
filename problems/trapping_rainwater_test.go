package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type rainwaterCase struct {
	Input  []int
	Answer int
}

func TestTrappedRainwater(t *testing.T) {
	for _, testCase := range []rainwaterCase{
		{
			Input:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			Answer: 6,
		},
		{
			Input:  []int{4, 2, 0, 3, 2, 5},
			Answer: 9,
		},
		{
			Input:  []int{1},
			Answer: 0,
		},
		{
			Input:  []int{2, 1, 2},
			Answer: 1,
		},
		{
			Input:  []int{1, 0, 0},
			Answer: 0,
		},
	} {
		result := problems.TrappedRainwater(testCase.Input)
		if result != testCase.Answer {
			t.Errorf("Expected %d, got %d\n", testCase.Answer, result)
		}
	}
}
