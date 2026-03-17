package daily

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	stack := []*ListNode{}
	curr := head
	for curr != nil {
		for len(stack) > 0 && stack[len(stack)-1].Val < curr.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, curr)
		curr = curr.Next
	}

	newHead := stack[0]
	currNode := newHead
	for i := 1; i < len(stack); i++ {
		currNode.Next = stack[i]
		currNode = currNode.Next
	}
	currNode.Next = nil
	return newHead
}
