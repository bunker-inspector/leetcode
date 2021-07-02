package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type input struct {
	Points [][2]int
	K      int
}

type maxValueOfEquationCase struct {
	Input  input
	Answer int
}

func TestMaxValueOfEquation(t *testing.T) {
	for _, testCase := range []maxValueOfEquationCase{
		{
			Input:  input{Points: [][2]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}, K: 1},
			Answer: 4,
		},
		{
			Input:  input{Points: [][2]int{{0, 0}, {3, 0}, {9, 2}}, K: 3},
			Answer: 3,
		},
		{
			Input: input{Points: [][2]int{
				{-19, -12},
				{-13, -18},
				{-12, 18},
				{-11, -8},
				{-8, 2},
				{-7, 12},
				{-5, 16},
				{-3, 9},
				{1, -7},
				{5, -4},
				{6, -20},
				{10, 4},
				{16, 4},
				{19, -9},
				{20, 19},
			},
				K: 6},
			Answer: 35,
		},
	} {
		result := problems.MaxValueOfEquation(testCase.Input.Points, testCase.Input.K)
		if result != testCase.Answer {
			t.Errorf("Given input %+v, %d:\n Expected %d, got %d.\n",
				testCase.Input.Points,
				testCase.Input.K,
				testCase.Answer,
				result,
			)
		}
	}
}
