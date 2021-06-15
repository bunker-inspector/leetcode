package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := ListNode{Val: vals[0]}
	curr := &head
	for _, val := range vals[1:] {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return &head
}

func (l *ListNode) Equals(other *ListNode) bool {
	thisPtr := l
	otherPtr := other
	if (thisPtr != nil && otherPtr == nil) || (thisPtr == nil && otherPtr != nil) {
		return false
	}

	for thisPtr != nil {
		if thisPtr.Val != otherPtr.Val {
			return false
		}
		thisPtr = thisPtr.Next
		otherPtr = otherPtr.Next
	}
	return true
}
