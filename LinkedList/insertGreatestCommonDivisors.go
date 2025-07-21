//2807. Insert Greatest Common Divisors in Linked List

package linkedList

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	pre := head.Next
	point := head
	for pre != nil {
		temp := new(ListNode)
		a, b := 0, 0
		a = max(point.Val, pre.Val)
		b = min(point.Val, pre.Val)
		temp.Val = gcd(a, b)
		point.Next = temp
		temp.Next = pre
		point = pre
		pre = pre.Next
	}
	return head
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

//2ms Beats 50%
//8.50MB Beats 15.29%
