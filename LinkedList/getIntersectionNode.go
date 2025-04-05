//160. Intersection of Two Linked Lists

package linkedList

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	count := 0
	for headA.Next != nil && headB.Next != nil {
		if headA == headB {
			return headA
		}
		if count%2 == 0 {
			headA = headA.Next
			if headA == headB {
				return headA
			} else {
				headA = headA.Next
			}
		}
		if count%2 == 1 {
			headB = headB.Next
			if headB == headA {
				return headB
			} else {
				headB = headB.Next
			}
		}
		count++
	}
	return nil
}
