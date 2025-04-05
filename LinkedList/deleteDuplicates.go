//83. Remove Duplicates from Sorted List

package linkedList

func deleteDuplicates(head *ListNode) *ListNode {
	clone := &ListNode{Val: 0}
	res := &ListNode{Val: 0}
	store := make([]int, 0, 300)
	i := 0
	for head != nil {
		if len(store) == 0 {
			store = append(store, head.Val)
			clone.Next = head
			clone = clone.Next
			res.Next = clone
			head = head.Next
			continue
		}
		if store[i] == head.Val {
			if head.Next == nil {
				clone.Next = nil
				break
			}
			head = head.Next
			continue
		}
		if store[i] != head.Val {
			store = append(store, head.Val)
			clone.Next = head
			clone = clone.Next
			head = head.Next
			i++
		}
	}
	return res.Next
}

//Runtime 0 ms Beats 100.00%
//Memory 4.97 MB Beats 27.28%
