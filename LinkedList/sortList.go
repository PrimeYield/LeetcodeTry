//148. Sort List

package linkedList

func sortList(head *ListNode) *ListNode {
	return split(head, nil)
}

func split(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return mergeAndSort(split(head, mid), split(mid, tail))
}

func mergeAndSort(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp, temp1, temp2 := dummy, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val > temp2.Val {
			temp.Next = temp2
			temp2 = temp2.Next
		} else {
			temp.Next = temp1
			temp1 = temp1.Next
		}
		temp = temp.Next
	}
	//處理殘留
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummy.Next
}

//runtime 3ms Beats 98.14%
//Memory 9.38MB Beats 53.13%
