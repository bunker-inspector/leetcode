package problems

func RotateList(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	// case of length 2
	if head.Next.Next == nil {
		if k%2 == 0 {
			return head
		} else {
			newHead := head.Next
			head.Next = nil
			newHead.Next = head
			return newHead
		}
	}

	ln := 1
	tailFind := head
	for tailFind.Next != nil {
		tailFind = tailFind.Next
		ln++
	}
	if k >= ln {
		k = k % ln
	}
	if k == 0 {
		return head
	}
	toSeek := ln - k - 1
	newTail := head
	for i := 0; i < toSeek; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	tailFind.Next = head
	newTail.Next = nil
	return newHead
}
