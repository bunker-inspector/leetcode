package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type interleavingStringInput struct {
	String1 string
	String2 string
	String3 string
}

type interleavingStringCase struct {
	Input  interleavingStringInput
	Answer bool
}

func TestIsInterleavingString(t *testing.T) {
	for _, testInput := range []interleavingStringCase{
		interleavingStringCase{
			Input: interleavingStringInput{
				String1: "aabcc",
				String2: "dbbca",
				String3: "aadbbcbcac",
			},
			Answer: true,
		},
		interleavingStringCase{
			Input: interleavingStringInput{
				String1: "aabcc",
				String2: "dbbca",
				String3: "aadbbbaccc",
			},
			Answer: false,
		},
		interleavingStringCase{
			Input: interleavingStringInput{
				"",
				"",
				"",
			},
			Answer: true,
		},
	} {
		result := problems.IsInterleavingString(testInput.Input.String1, testInput.Input.String2, testInput.Input.String3)
		if result != testInput.Answer {
			t.Errorf("Given input %s, %s, %s:\n Expected %t, got %t.\n",
				testInput.Input.String1,
				testInput.Input.String2,
				testInput.Input.String3,
				testInput.Answer,
				result)
		}
	}
}
