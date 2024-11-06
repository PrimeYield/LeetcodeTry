package binaryTree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if p.Val == q.Val && p.Left == nil && q.Left == nil && p.Right == nil && q.Right == nil {
		return true
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
