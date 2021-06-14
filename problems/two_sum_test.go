package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type twoSumInput struct {
	Nums   []int
	Target int
}

type twoSumCase struct {
	Input  twoSumInput
	Answer []int
}

func TestTwoSum(t *testing.T) {
	for _, testInput := range []twoSumCase{
		twoSumCase{
			Input:  twoSumInput{Nums: []int{2, 7, 11, 15}, Target: 9},
			Answer: []int{0, 1},
		},
		twoSumCase{
			Input:  twoSumInput{Nums: []int{3, 2, 4}, Target: 6},
			Answer: []int{1, 2},
		},
		twoSumCase{
			Input:  twoSumInput{Nums: []int{3, 3}, Target: 6},
			Answer: []int{0, 1},
		},
	} {
		result := problems.TwoSum(testInput.Input.Nums, testInput.Input.Target)
		if result[0] != testInput.Answer[0] && result[1] != testInput.Answer[1] {
			t.Errorf("Given input %+v, %d:\n Expected %+v, got %+v.\n",
				testInput.Input.Nums,
				testInput.Input.Target,
				testInput.Answer,
				result)
		}
	}
}
