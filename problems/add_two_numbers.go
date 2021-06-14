package problems

import (
	"strconv"
	"strings"
)

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

func (l *ListNode) String() string {
	var sb strings.Builder
	cursor := l
	if cursor != nil {
		sb.WriteString(strconv.FormatInt(int64(cursor.Val), 10))
		cursor = cursor.Next
	} else {
		return ""
	}
	for cursor != nil {
		sb.WriteString(", ")
		sb.WriteString(strconv.FormatInt(int64(cursor.Val), 10))
		cursor = cursor.Next
	}
	return sb.String()
}

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