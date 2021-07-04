package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type findInRotatedArrayInput struct {
	Nums   []int
	Target int
}

type findInRotatedArrayCase struct {
	Input  findInRotatedArrayInput
	Answer int
}

func TestFindInRotatedArray(t *testing.T) {
	for _, testCase := range []findInRotatedArrayCase{
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{4, 5, 6, 7, 0, 1, 2},
				Target: 0,
			},
			Answer: 4,
		},
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{4, 5, 6, 7, 0, 1, 2},
				Target: 3,
			},
			Answer: -1,
		},
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{1},
				Target: 0,
			},
			Answer: -1,
		},
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{6, 7, 8, 0, 1, 2, 3, 4, 5},
				Target: 5,
			},
			Answer: 8,
		},
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{6, 7, 8, 0, 1, 2, 3, 4, 5},
				Target: 6,
			},
			Answer: 0,
		},
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{},
				Target: 5,
			},
			Answer: -1,
		},
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{1, 2, 3, 4},
				Target: 3,
			},
			Answer: 2,
		},
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{1, 3},
				Target: 3,
			},
			Answer: 1,
		},
		{
			Input: findInRotatedArrayInput{
				Nums:   []int{1, 2, 3},
				Target: 3,
			},
			Answer: 2,
		},
	} {
		result := problems.FindInRotatedArray(testCase.Input.Nums, testCase.Input.Target)
		if result != testCase.Answer {
			t.Errorf("Expected %d, got %d\n", testCase.Answer, result)
			t.Fail()
		}
	}
}
