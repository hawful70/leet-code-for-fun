package daily

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// count length
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	dummy := &ListNode{Val: 0, Next: head}
	groupPrev := dummy
	curr = head

	for i := 0; i < length/k; i++ {
		groupTail := curr  // curr is the future tail of this reversed group
		var prev *ListNode // nil in Go, equivalent to Python's None

		for j := 0; j < k; j++ {
			nxt := curr.Next
			curr.Next = prev
			prev = curr
			curr = nxt
		}

		groupPrev.Next = prev // stitch left side
		groupTail.Next = curr // stitch right side
		groupPrev = groupTail // slide boundary forward
	}

	return dummy.Next
}
