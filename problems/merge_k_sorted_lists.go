package problems

import (
	"container/heap"
)

type NodeHeap []*ListNode

func (h NodeHeap) Len() int { return len(h) }
func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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
