package daily

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	dfsDiameterOfBinaryTree(root)
	return maxDiameter
}

func dfsDiameterOfBinaryTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := dfsDiameterOfBinaryTree(node.Left)
	right := dfsDiameterOfBinaryTree(node.Right)
	maxDiameter = max(maxDiameter, left+right)
	return 1 + max(left, right)
}
