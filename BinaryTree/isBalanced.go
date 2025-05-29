//110. Balanced Binary Tree

package binaryTree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	return (leftHeight-rightHeight == 1 || leftHeight-rightHeight == 0 || leftHeight-rightHeight == -1) && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

//runtime 0ms Memory 7.35MB
