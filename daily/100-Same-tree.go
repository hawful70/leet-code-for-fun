package daily

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return dfsTree(p, q)
}

func dfsTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	left := dfsTree(p.Left, q.Left)
	right := dfsTree(p.Right, q.Right)
	return left && right
}
