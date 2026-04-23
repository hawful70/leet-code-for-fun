package daily

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	curr := head

	for curr != nil {
		next_node := curr.Next
		prev := dummy

		for prev.Next != nil && prev.Next.Val < curr.Val {
			prev = prev.Next
		}

		curr.Next = prev.Next
		prev.Next = curr

		curr = next_node
	}

	return dummy.Next
}
