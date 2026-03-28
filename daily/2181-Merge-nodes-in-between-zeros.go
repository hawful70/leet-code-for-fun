package daily

func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	curr := dummy

	zeroCount := 0
	sumNodes := 0
	for head != nil {
		if head.Val == 0 {
			zeroCount++
		}
		sumNodes += head.Val

		if zeroCount >= 2 {
			newNode := &ListNode{Val: sumNodes}
			if sumNodes > 0 {
				curr.Next = newNode
				curr = curr.Next
			}
			zeroCount = 1
			sumNodes = 0
		}

		head = head.Next
	}
	return dummy.Next
}
