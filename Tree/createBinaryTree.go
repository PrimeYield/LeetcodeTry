//2196. Create Binary Tree From Descriptions

package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*TreeNode)
	for i := 0; i < len(descriptions); i++ {
		nodes[descriptions[i][0]] = &TreeNode{Val: descriptions[i][0]}
	}
	res := &TreeNode{}
	//root是從來沒有在description[i][1]出現過的description[i][0]
	isChild := make(map[int]bool)
	// isChild := make([]bool, len(descriptions), len(descriptions))
	for i := 0; i < len(descriptions); i++ {
		isChild[descriptions[i][1]] = true
	}
	for k, root := range nodes {
		if _, exist := isChild[k]; !exist {
			res = root
		}
	}
	for i := 0; i < len(descriptions); i++ {
		if _, exist := nodes[descriptions[i][1]]; exist {
			//代表這個node有孩子，同時也有父母
			if descriptions[i][2] == 1 {
				nodes[descriptions[i][0]].Left = nodes[descriptions[i][1]]
			} else {
				nodes[descriptions[i][0]].Right = nodes[descriptions[i][1]]
			}
			// res = nodes[descriptions[i][0]]
		} else if !exist {
			//代表這個node沒有孩子，是葉子
			// nodes[descriptions[i][0]].Left OR .Right
			if descriptions[i][2] == 1 {
				nodes[descriptions[i][0]].Left = &TreeNode{Val: descriptions[i][1]}
			} else {
				nodes[descriptions[i][0]].Right = &TreeNode{Val: descriptions[i][1]}
			}
		}
	}
	return res
}
