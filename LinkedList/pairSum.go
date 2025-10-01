//2130. Maximum Twin Sum of a Linked List

package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	stack := []int{}
	fast := head
	slow := head
	for fast != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	res := 0
	for slow != nil {
		sum := slow.Val + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		slow = slow.Next
		if sum > res {
			res = sum
		}
	}
	return res
}
