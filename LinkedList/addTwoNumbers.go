package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headNum := &ListNode{Val: 0} //初始化根節點&ListNode{Val: 0},虛擬根節點
	curList := headNum           //游標,初始化指向headNum,表示從虛擬頭節點開始構建新的鏈結串列。
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 { //進位!=0在避免忽略最後兩個值相加後遺留的進位
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}
		sum := carry + x + y
		carry = sum / 10 //用來存放進位
		curList.Next = &ListNode{Val: sum % 10}
		curList = curList.Next //前進
		if l1 != nil {
			l1 = l1.Next //前進
		}
		if l2 != nil {
			l2 = l2.Next //前進
		}
	}
	return headNum.Next //headNum 本身的值為 0 不需要
}
