package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type largestRectangleCase struct {
	Input  []int
	Answer int
}

func TestLargestRectangle(t *testing.T) {
	for _, testCase := range []largestRectangleCase{
		{
			Input:  []int{2, 1, 5, 6, 2, 3},
			Answer: 10,
		},
		{
			Input:  []int{2, 4},
			Answer: 4,
		},
		{
			Input:  []int{},
			Answer: 0,
		},
		{
			Input:  []int{3, 3, 3},
			Answer: 9,
		},
		{
			Input:  []int{5, 4, 3, 2, 1},
			Answer: 9,
		},
		{
			Input:  []int{1, 2, 3, 4, 5},
			Answer: 9,
		},
		{
			Input:  []int{3},
			Answer: 3,
		},
	} {
		result := problems.LargestRectangle(testCase.Input)
		if result != testCase.Answer {
			t.Errorf("Expected %d, but got %d\n", testCase.Answer, result)
			t.Fail()
		}
	}
}
