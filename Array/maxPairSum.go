// 1877. Minimize Maximum Pair Sum in Array
package array

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	var max int

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		var sum int
		sum = nums[i] + nums[j]
		if max < sum {
			max = sum
		}
	}
	return max
}
