package problems

import (
	"container/heap"
)

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := NodeHeap{}
	for _, list := range lists {
		if list != nil {
			h = append(h, list)
		}
	}
	heap.Init(&h)

	if h.Len() == 0 {
		return nil
	}

	head := heap.Pop(&h).(*ListNode)
	if head.Next != nil {
		heap.Push(&h, head.Next)
		head.Next = nil
	}
	out := head
	tail := out
	for h.Len() > 0 {
		head := heap.Pop(&h).(*ListNode)
		if head.Next != nil {
			heap.Push(&h, head.Next)
			head.Next = nil
		}
		tail.Next = head
		tail = tail.Next
	}
	return out
}
