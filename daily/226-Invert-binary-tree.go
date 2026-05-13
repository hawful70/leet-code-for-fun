package daily

func invertTree(root *TreeNode) *TreeNode {
	return dfsInvertTree(root)
}

func dfsInvertTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left, node.Right = node.Right, node.Left
	dfsInvertTree(node.Left)
	dfsInvertTree(node.Right)
	return node
}
