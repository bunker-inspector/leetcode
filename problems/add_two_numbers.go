package problems

func handleCarry(l *ListNode, carry *int) {
	*carry = 0
	if l.Val >= 10 {
		*carry = 1
		l.Val -= 10
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result *ListNode
		carry  int
	)
	result = &ListNode{}
	out := result
	result.Val = l1.Val + l2.Val
	l1 = l1.Next
	l2 = l2.Next

	handleCarry(result, &carry)
	for l1 != nil && l2 != nil {
		result.Next = &ListNode{Val: l1.Val + l2.Val + carry}
		result = result.Next
		handleCarry(result, &carry)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		result.Next = &ListNode{Val: l1.Val + carry}
		result = result.Next
		handleCarry(result, &carry)
		l1 = l1.Next
	}
	for l2 != nil {
		result.Next = &ListNode{Val: l2.Val + carry}
		result = result.Next
		handleCarry(result, &carry)
		l2 = l2.Next
	}
	if carry == 1 {
		result.Next = &ListNode{Val: 1}
	}

	return out
}
