//617. Merge Two Binary Trees

package binaryTree

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return root1
	}
	if root1 == nil && root2 != nil {
		root1 = root2
		return root1
	}
	if root1 != nil && root2 == nil {
		//do nothing
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

//runtime 0ms
//Memory 9.39MB
