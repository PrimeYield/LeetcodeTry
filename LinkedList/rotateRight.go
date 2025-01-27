//61. Rotate List

package linkedList

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	var count int = 1
	cur := head
	for cur.Next != nil {
		count++
		cur = cur.Next
	}
	cur.Next = head
	k = k % count
	count = count - k
	for i := 0; i < count; i++ {
		cur = cur.Next
	}
	iter := cur.Next
	cur.Next = nil
	return iter
}

/*
Runtime 0ms Beats 100.00%
Memory 4.50MB Beats 11.26%
*/
