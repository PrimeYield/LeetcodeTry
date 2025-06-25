// 206. Reverse Linked List
package linkedList

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy *ListNode = nil
	// dummy.Next = head
	back := head
	front := head.Next
	for front != nil {
		back.Next = dummy
		dummy = back
		back = front
		front = front.Next
	}
	back.Next = dummy
	dummy = back
	return dummy
}

//runtime 0ms
//Memory 4.44MB
