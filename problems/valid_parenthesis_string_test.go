package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type validParenthesisCase struct {
	Input  string
	Answer bool
}

func TestValidParenthesis(t *testing.T) {
	for _, testCase := range []validParenthesisCase{
		validParenthesisCase{
			Input:  "()",
			Answer: true,
		},
		validParenthesisCase{
			Input:  "(*)",
			Answer: true,
		},
		validParenthesisCase{
			Input:  "()*)",
			Answer: true,
		},
		validParenthesisCase{
			Input:  "(*))",
			Answer: true,
		},
		validParenthesisCase{
			Input:  "()",
			Answer: true,
		},
		validParenthesisCase{
			Input:  "))",
			Answer: false,
		},
		validParenthesisCase{
			Input:  "**))",
			Answer: true,
		},
		validParenthesisCase{
			Input:  "*",
			Answer: true,
		},
		validParenthesisCase{
			Input:  "*(",
			Answer: false,
		},
		validParenthesisCase{
			Input:  "(",
			Answer: false,
		},
	} {
		result := problems.ValidParenthesis(testCase.Input)
		if result != testCase.Answer {
			t.Errorf("Given input %s:\n Expected %t, got %t.\n",
				testCase.Input,
				testCase.Answer,
				result,
			)
		}
	}
}
