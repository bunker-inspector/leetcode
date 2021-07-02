package problems_test

import (
	"testing"

	"github.com/bunker-inspector/leetcode-go/problems"
)

type mergeKListsCase struct {
	Input  []*problems.ListNode
	Answer *problems.ListNode
}

func TestMergeKLists(t *testing.T) {
	for _, testCase := range []mergeKListsCase{
		{
			Input: []*problems.ListNode{
				problems.CreateList(1, 4, 5),
				problems.CreateList(1, 3, 4),
				problems.CreateList(2, 6),
			},
			Answer: problems.CreateList(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			Input:  []*problems.ListNode{},
			Answer: problems.CreateList(),
		},
		{
			Input: []*problems.ListNode{
				problems.CreateList(),
			},
			Answer: problems.CreateList(),
		},
		{
			Input: []*problems.ListNode{
				problems.CreateList(1, 2, 3, 4, 5, 6),
				problems.CreateList(7),
				problems.CreateList(),
			},
			Answer: problems.CreateList(1, 2, 3, 4, 5, 6, 7),
		},
		{
			Input: []*problems.ListNode{
				problems.CreateList(1, 2, 3, 4, 5, 6, 8, 9),
				problems.CreateList(7),
				problems.CreateList(),
			},
			Answer: problems.CreateList(1, 2, 3, 4, 5, 6, 7, 8, 9),
		},
	} {
		result := problems.MergeKLists(testCase.Input)
		if !result.Equals(testCase.Answer) {
			t.Errorf("Expected %s, got %s\n", testCase.Answer, result)
			t.Fail()
		}
	}
}
