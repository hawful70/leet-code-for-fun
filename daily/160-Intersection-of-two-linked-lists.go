package daily

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	seen := make(map[*ListNode]bool)

	for headA != nil {
		seen[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if seen[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
