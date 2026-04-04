package daily

func swapNodes(head *ListNode, k int) *ListNode {
	left := head
	for i := 0; i < k-1; i++ {
		left = left.Next
	}

	fast := left
	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	left.Val, slow.Val = slow.Val, left.Val
	return head
}
