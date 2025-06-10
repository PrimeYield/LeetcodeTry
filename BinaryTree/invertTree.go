//226. Invert Binary Tree

package binaryTree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dummy := root
	newLeft := root.Right
	newRight := root.Left
	root.Left = newLeft
	root.Right = newRight
	invertTree(root.Left)
	invertTree(root.Right)
	return dummy
}

//runtime 0ms
//Memory 4.14MB Beats 53.09%
