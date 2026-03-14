package daily

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{0, head}
	before := dummy

	for i := 0; i < left-1; i++ {
		before = before.Next
	}

	curr := before.Next
	var prev *ListNode
	for i := 0; i < right-left+1; i++ {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}
	before.Next.Next = curr
	before.Next = prev
	return dummy.Next
}
