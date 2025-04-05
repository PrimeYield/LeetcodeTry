// 141. Linked List Cycle
package linkedList

/*
func hasCycle(head *ListNode) bool {
	count := 0
	for head.Next != nil {
		count++
		head = head.Next
		if count > 10000 {
			break
		}
	}
	return count > 1
}

Runtime 2 m2 Beats 96.57%
Memory 6.14 MB Beats 93.29%

做法有點利用題目的設計漏洞
*/

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

/*
Runtime 8 ms Beats 36.05%
Memory 6.15MB Beats 93.29%
*/

/*犯錯
for slow.Next != nil || fast.Next != nil 的條件會造成 [1,2] pos = -1的情況下
fast.Next.Next會指在nil值上，但slow.Next != nil 依然成立
會使得fast.Next.Next指向空的記憶體位址

for fast.Next != nil 的條件，則是fast指針指在空的記憶體位址，fast這時並不是nil，而是沒有宣告(?)

因此條件式需要為 fast != nil && fast.Next != nil
*/
