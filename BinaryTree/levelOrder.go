//102. Binary Tree Level Order Traversal

package binaryTree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		var res [][]int
		res = append(res, []int{root.Val})
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	count := 0
	curCount := 1 //第一層
	var res [][]int
	for {
		temp := make([]int, 0)
		for i := 0; i < curCount; i++ {
			if queue[i].Left != nil && queue[i].Right != nil {
				count += 2
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
			} else if queue[i].Left == nil && queue[i].Right != nil {
				count += 1
				queue = append(queue, queue[i].Right)
			} else if queue[i].Left != nil && queue[i].Right == nil {
				count += 1
				queue = append(queue, queue[i].Left)
			} else if queue[i].Left == nil && queue[i].Right == nil {
			}
			temp = append(temp, queue[i].Val)
		}
		queue = queue[curCount:]
		if len(temp) == 0 {
			break
		}
		res = append(res, temp)
		curCount = count
		count = 0
	}
	return res
}

//runtime 0ms
//Memory 5.47MB Beats 96.06%
