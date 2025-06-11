//814. Binary Tree Pruning

package binaryTree

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		root = nil
	}
	// if root.Left != nil || root.Right != nil {
	// 	return root
	// }
	return root
}

//runtime 0ms
//Memory 4.31MB
