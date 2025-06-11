//543. Diameter of Binary Tree

package binaryTree

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := dfs(root.Left)
		rightHeight := dfs(root.Right)
		maxDiameter = max(maxDiameter, leftHeight+rightHeight)

		return max(leftHeight, rightHeight) + 1
	}
	dfs(root)
	return maxDiameter
}

//runtime 0ms
//Memory 8.61MB
