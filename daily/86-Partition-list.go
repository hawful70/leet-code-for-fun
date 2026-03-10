package daily

func partition(head *ListNode, x int) *ListNode {
	beforeDummy := &ListNode{0, nil}
	afterDummy := &ListNode{0, nil}

	before := beforeDummy
	after := afterDummy

	curr := head
	for curr != nil {
		if curr.Val < x {
			before.Next = curr
			before = before.Next
		} else {
			after.Next = curr
			after = after.Next
		}
		curr = curr.Next
	}

	after.Next = nil
	before.Next = afterDummy.Next
	return beforeDummy.Next
}
