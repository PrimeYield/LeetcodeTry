//2574. Left and Right Sum Differences

package array

import "fmt"

func leftRightDifference(nums []int) []int {
	var sum int
	length := len(nums)
	left := make([]int, length, length)
	right := make([]int, length, length)
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		left[i] = sum
		sum = sum + nums[i]
	}
	sum = 0
	for i := length - 1; i >= 0; i-- {
		right[i] = sum
		sum = sum + nums[i]
	}
	sum = 0
	for i := 0; i < length; i++ {
		sum = left[i] - right[i]
		if sum < 0 {
			sum = -sum
		}
		res[i] = sum
	}
	fmt.Println(left)
	fmt.Println(right)
	return res
}
