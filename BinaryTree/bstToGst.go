//1038. Binary Search Tree to Greater Sum Tree

package binaryTree

func bstToGst(root *TreeNode) *TreeNode {
	// if root == nil {
	// 	return nil
	// }
	// root.Right = bstToGst(root.Right)
	sum := 0
	dfs(root.Right, &sum)
	sum += root.Val
	root.Val = sum
	dfs(root.Left, &sum)
	return root
}

func dfs(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	dfs(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	dfs(root.Left, sum)
	return root
}

//0ms
//4.26MB 50%
