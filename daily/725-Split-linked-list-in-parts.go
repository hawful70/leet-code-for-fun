package daily

func splitListToParts(head *ListNode, k int) []*ListNode {
	ans := make([]*ListNode, k)
	size := 0
	current := head
	for current != nil {
		size++
		current = current.Next
	}

	splitSize := size / k
	remainSize := size % k
	current = head
	for i := 0; i < k; i++ {
		ans[i] = current
		currentSize := splitSize
		if remainSize > 0 {
			currentSize++
			remainSize--
		}

		for j := 0; j < currentSize-1; j++ {
			if current != nil {
				current = current.Next
			}
		}
		if current != nil {
			nxt := current.Next
			current.Next = nil
			current = nxt
		}
	}

	return ans
}
