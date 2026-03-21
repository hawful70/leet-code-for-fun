package daily

func nextLargerNodes(head *ListNode) []int {
	// Convert linked list to slice for easier indexing
	var nodes []*ListNode
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	n := len(nodes)
	result := make([]int, n)

	// Stack stores indices of nodes in decreasing order of value
	// Format: [index, value]
	stack := [][2]int{}

	for i := 0; i < n; i++ {
		// While stack is not empty and current node is greater than stack top
		for len(stack) > 0 && nodes[i].Val > stack[len(stack)-1][1] {
			// Pop from stack and set the result for that index
			topIndex := stack[len(stack)-1][0]
			stack = stack[:len(stack)-1]
			result[topIndex] = nodes[i].Val
		}
		// Push current node's index and value onto stack
		stack = append(stack, [2]int{i, nodes[i].Val})
	}

	// Remaining elements in stack have no greater node to their right
	// Their result values are already 0 (default for int slice)

	return result
}
