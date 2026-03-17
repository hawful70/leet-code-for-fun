package daily

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, list1}
	before := dummy

	for i := 0; i < a; i++ {
		before = before.Next
	}

	curr := before.Next
	for i := 0; i < b-a+1; i++ {
		curr = curr.Next
	}

	list2Curr := list2
	for list2Curr.Next != nil {
		list2Curr = list2Curr.Next
	}

	list2Curr.Next = curr
	before.Next = list2
	return dummy.Next
}
