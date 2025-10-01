//23. Merge k Sorted Lists

package linkedList

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		res := &ListNode{}
		dummy := res
		for lists[0] != nil && lists[1] != nil {
			if lists[0].Val > lists[1].Val {
				dummy.Next = lists[1]
				dummy = dummy.Next
				lists[1] = lists[1].Next
			} else {
				dummy.Next = lists[0]
				dummy = dummy.Next
				lists[0] = lists[0].Next
			}
		}
		if lists[0] != nil {
			dummy.Next = lists[0]
		} else {
			dummy.Next = lists[1]
		}
		return res.Next
	}
	mid := len(lists) / 2
	return mergeKLists([]*ListNode{mergeKLists(lists[:mid]), mergeKLists(lists[mid:])})
}

//0ms 100%
//6.43MB 97.90%
