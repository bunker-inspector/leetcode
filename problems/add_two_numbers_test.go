package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type AddTwoNumbersInput struct {
	List1  *problems.ListNode
	List2  *problems.ListNode
	Answer *problems.ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	for _, testInput := range []AddTwoNumbersInput{
		AddTwoNumbersInput{
			List1:  problems.CreateList(2, 4, 3),
			List2:  problems.CreateList(5, 6, 4),
			Answer: problems.CreateList(7, 0, 8),
		},
		AddTwoNumbersInput{
			List1:  problems.CreateList(9, 9, 9, 9, 9, 9, 9),
			List2:  problems.CreateList(9, 9, 9, 9),
			Answer: problems.CreateList(8, 9, 9, 9, 0, 0, 0, 1),
		},
		AddTwoNumbersInput{
			List1:  problems.CreateList(0),
			List2:  problems.CreateList(0),
			Answer: problems.CreateList(0),
		},
	} {
		result := problems.AddTwoNumbers(testInput.List1, testInput.List2)
		resCurr := result
		ansCurr := testInput.Answer
		for ansCurr != nil {
			if resCurr.Val != ansCurr.Val {
				t.Errorf("Result did not match expected!: %+v -- %+v", result, testInput.Answer)
				t.Fail()
				break
			}
			resCurr = resCurr.Next
			ansCurr = ansCurr.Next
		}
	}
}
