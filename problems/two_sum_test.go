package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type input struct {
	Nums   []int
	Target int
}

type TwoSumInput struct {
	Input  input
	Answer []int
}

func TestTwoSum(t *testing.T) {
	for _, testInput := range []TwoSumInput{
		TwoSumInput{
			Input:  input{Nums: []int{2, 7, 11, 15}, Target: 9},
			Answer: []int{0, 1},
		},
		TwoSumInput{
			Input:  input{Nums: []int{3, 2, 4}, Target: 6},
			Answer: []int{1, 2},
		},
		TwoSumInput{
			Input:  input{Nums: []int{3, 3}, Target: 6},
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
