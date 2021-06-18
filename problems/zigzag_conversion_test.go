package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type zigZagInput struct {
	S    string
	Rows int
}

type ZigZagConversionCase struct {
	Input  zigZagInput
	Answer string
}

func TestZigZagConversion(t *testing.T) {
	for _, testCase := range []ZigZagConversionCase{
		ZigZagConversionCase{
			Input: zigZagInput{
				S:    "PAYPALISHIRING",
				Rows: 3,
			},
			Answer: "PAHNAPLSIIGYIR",
		},
		ZigZagConversionCase{
			Input: zigZagInput{
				S:    "PAYPALISHIRING",
				Rows: 4,
			},
			Answer: "PINALSIGYAHRPI",
		},
		ZigZagConversionCase{
			Input: zigZagInput{
				S:    "A",
				Rows: 1,
			},
			Answer: "A",
		},
		ZigZagConversionCase{
			Input: zigZagInput{
				S:    "A",
				Rows: 0,
			},
			Answer: "",
		},
		ZigZagConversionCase{
			Input: zigZagInput{
				S:    "ABCD",
				Rows: 5,
			},
			Answer: "ABCD",
		},
		ZigZagConversionCase{
			Input: zigZagInput{
				S:    "",
				Rows: 8,
			},
			Answer: "",
		},
	} {
		result := problems.ConvertZigZag(testCase.Input.S, testCase.Input.Rows)
		if result != testCase.Answer {
			t.Errorf("Given input %s, %d: \n Expected %s, got %s.\n",
				testCase.Input.S,
				testCase.Input.Rows,
				testCase.Answer,
				result)
		}
	}
}
