package daily

func preorderTraversal(root *TreeNode) []int {
	var results []int
	dfsPreorder(root, &results)
	return results
}

func dfsPreorder(node *TreeNode, results *[]int) {
	if node == nil {
		return
	}
	*results = append(*results, node.Val)
	dfsPreorder(node.Left, results)
	dfsPreorder(node.Right, results)
}
