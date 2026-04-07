package daily

func pairSum(head *ListNode) int {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	maxTwinSum := 0
	left := head
	right := prev
	for right != nil {
		currentTwinSum := left.Val + right.Val
		if currentTwinSum > maxTwinSum {
			maxTwinSum = currentTwinSum
		}
		left = left.Next
		right = right.Next
	}

	return maxTwinSum
}
