package problems

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

type MaxIntHeap []int

func (h MaxIntHeap) Len() int            { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MaxIntHeap) Peek() int { return (*h)[0] }

type MinIntHeap []int

func (h MinIntHeap) Len() int            { return len(h) }
func (h MinIntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MinIntHeap) Peek() int { return (*h)[0] }
