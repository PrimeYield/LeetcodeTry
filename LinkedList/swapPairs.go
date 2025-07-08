//24. Swap Nodes in Pairs

package linkedList

func swapPairs(head *ListNode) *ListNode {
	dummy := head
	res := &ListNode{}
	curr := res
	for dummy != nil && dummy.Next != nil {
		tempNext := &ListNode{Val: dummy.Val}
		temp := &ListNode{Val: dummy.Next.Val}
		temp.Next = tempNext
		curr.Next = temp
		curr = curr.Next.Next
		dummy = dummy.Next.Next
	}
	if dummy != nil {
		curr.Next = dummy
	}
	return res.Next
}

//0ms
//Memory 3.96% Beats 87.43%
