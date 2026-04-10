package daily

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	l1Stacks := []int{}
	l2Stacks := []int{}

	for l1 != nil {
		l1Stacks = append(l1Stacks, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		l2Stacks = append(l2Stacks, l2.Val)
		l2 = l2.Next
	}

	carry := 0
	var head *ListNode

	for len(l1Stacks) > 0 || len(l2Stacks) > 0 || carry > 0 {
		item1 := 0
		if len(l1Stacks) > 0 {
			item1 = l1Stacks[len(l1Stacks)-1]
			l1Stacks = l1Stacks[:len(l1Stacks)-1]
		}

		item2 := 0
		if len(l2Stacks) > 0 {
			item2 = l2Stacks[len(l2Stacks)-1]
			l2Stacks = l2Stacks[:len(l2Stacks)-1]
		}

		total := item1 + item2 + carry
		carry = total / 10
		newDigit := total % 10

		newNode := &ListNode{Val: newDigit}
		newNode.Next = head
		head = newNode
	}

	return head
}
