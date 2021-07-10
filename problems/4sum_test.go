package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type fourSumInput struct {
	Nums   []int
	Target int
}

type FourSumCase struct {
	Input  fourSumInput
	Answer problems.IntArray4Set
}

func TestFourSum(t *testing.T) {
	for _, testCase := range []FourSumCase{
		{
			Input: fourSumInput{
				Nums:   []int{1, 0, -1, 0, -2, 2},
				Target: 0,
			},
			Answer: *problems.NewIntArray4Set([4]int{-2, -1, 1, 2}, [4]int{-2, 0, 0, 2}, [4]int{-1, 0, 0, 1}),
		},
		{
			Input: fourSumInput{
				Nums:   []int{2, 2, 2, 2},
				Target: 8,
			},
			Answer: *problems.NewIntArray4Set([4]int{2, 2, 2, 2}),
		},
	} {
		result := problems.FourSum(testCase.Input.Nums, testCase.Input.Target)
		if len(result) != len(testCase.Answer) || !testCase.Answer.Contains(result...) {
			t.Errorf("Got %+v, expects %+v\n", result, testCase.Answer)
			t.Fail()
		}
	}
}
