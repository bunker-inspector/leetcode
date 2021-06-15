package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type rotateListCase struct {
	List   *problems.ListNode
	K      int
	Answer *problems.ListNode
}

func TestRotateList(t *testing.T) {
	for _, testInput := range []rotateListCase{
		rotateListCase{
			List:   problems.CreateList(1, 2, 3, 4, 5),
			K:      2,
			Answer: problems.CreateList(4, 5, 1, 2, 3),
		},
		rotateListCase{
			List:   problems.CreateList(0, 1, 2),
			K:      4,
			Answer: problems.CreateList(2, 0, 1),
		},
		rotateListCase{
			List:   &problems.ListNode{},
			K:      0,
			Answer: &problems.ListNode{},
		},
		rotateListCase{
			List:   problems.CreateList(1, 2),
			K:      0,
			Answer: problems.CreateList(1, 2),
		},
		rotateListCase{
			List:   problems.CreateList(1, 2, 3),
			K:      2000000000,
			Answer: problems.CreateList(2, 3, 1),
		},
		rotateListCase{
			List:   problems.CreateList(1, 2, 3),
			K:      3,
			Answer: problems.CreateList(1, 2, 3),
		},
	} {
		result := problems.RotateList(testInput.List, testInput.K)
		if !result.Equals(testInput.Answer) {
			t.Errorf("Result did not match expected!: %+v -- %+v", result, testInput.Answer)
			t.Fail()
		}
	}
}
