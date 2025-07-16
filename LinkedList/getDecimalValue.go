//1290. Convert Binary Number in a Linked List to Integer

package linkedList

import "math"

func getDecimalValue(head *ListNode) int {
	bin := make([]int, 0)
	for head != nil {
		bin = append(bin, head.Val)
		head = head.Next
	}
	res := 0
	for i := len(bin) - 1; i >= 0; i-- {
		res = res + (bin[i] * int(math.Pow(2, float64(len(bin)-1-i))))
	}
	return res
}

//0ms 4.17MB
