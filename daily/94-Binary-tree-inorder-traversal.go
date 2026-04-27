package daily

func inorderTraversal(root *TreeNode) []int {
	var results []int
	dfsInorder(root, &results)
	return results
}

func dfsInorder(node *TreeNode, results *[]int) {
	if node == nil {
		return
	}
	dfsInorder(node.Left, results)
	*results = append(*results, node.Val)
	dfsInorder(node.Right, results)
}
