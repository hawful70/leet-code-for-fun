package daily

func minDepth(root *TreeNode) int {
	return dfsMinDepth(root)
}

func dfsMinDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := dfsMinDepth(node.Left)
	right := dfsMinDepth(node.Right)
	if node.Left == nil {
		return 1 + right
	}
	if node.Right == nil {
		return 1 + left
	}
	return 1 + min(left, right)
}
