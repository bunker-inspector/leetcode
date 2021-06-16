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
		maxValueOfEquationCase{
			Input:  input{Points: [][2]int{[2]int{1, 3}, [2]int{2, 0}, [2]int{5, 10}, [2]int{6, -10}}, K: 1},
			Answer: 4,
		},
		maxValueOfEquationCase{
			Input:  input{Points: [][2]int{[2]int{0, 0}, [2]int{3, 0}, [2]int{9, 2}}, K: 3},
			Answer: 3,
		},
		maxValueOfEquationCase{
			Input: input{Points: [][2]int{
				[2]int{-19, -12},
				[2]int{-13, -18},
				[2]int{-12, 18},
				[2]int{-11, -8},
				[2]int{-8, 2},
				[2]int{-7, 12},
				[2]int{-5, 16},
				[2]int{-3, 9},
				[2]int{1, -7},
				[2]int{5, -4},
				[2]int{6, -20},
				[2]int{10, 4},
				[2]int{16, 4},
				[2]int{19, -9},
				[2]int{20, 19},
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
