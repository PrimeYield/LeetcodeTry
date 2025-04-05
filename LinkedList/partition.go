//86. Partition List

package linkedList

func partition(head *ListNode, x int) *ListNode {
	//思路：兩個鏈表結合 <x的一條 >=x的一條 <x .Next = >=x
	//[1,4,3,2,5,2], x = 3
	//1,2,2
	//4,3,5
	//if 1,2,2 ; 2.next == nil {merge 4,3,5}
	smallHead := &ListNode{}
	bigHead := &ListNode{}
	small := smallHead
	big := bigHead
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}

//Runtime 0ms Beats 100.00%
//Memory 4.30MB beats 59.38%
