package daily

func postorderTraversal(root *TreeNode) []int {
	var results []int
	dfs(root, &results)
	return results
}

func dfs(node *TreeNode, results *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, results)
	dfs(node.Right, results)
	*results = append(*results, node.Val)
}
