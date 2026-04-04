package daily

func modifiedList(nums []int, head *ListNode) *ListNode {
	seen := make(map[int]bool)
	for _, num := range nums {
		seen[num] = true
	}

	dummy := &ListNode{Val: 0}
	curr := dummy
	for head != nil {
		if !seen[head.Val] {
			curr.Next = head
			curr = curr.Next
		}
		head = head.Next
	}
	curr.Next = nil
	return dummy.Next
}
