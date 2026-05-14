package daily

func maxDepth(root *TreeNode) int {
	return dfsMaxDepth(root)
}

func dfsMaxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := dfsMaxDepth(node.Left)
	right := dfsMaxDepth(node.Right)
	return 1 + max(left, right)
}
