package daily

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{carry % 10, nil}
		curr = curr.Next
		carry /= 10
	}
	return dummy.Next
}
