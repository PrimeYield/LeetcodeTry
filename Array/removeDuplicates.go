package array

import "sort"

func removeDuplicates(nums []int) int {
	cur := nums[0]
	for i, num := range nums {
		if i > 0 && num == cur {
			nums[i] = 99999999
		} else {
			cur = num
		}
	}
	sort.Ints(nums)
	for i, num := range nums {
		if num < 99999999 {
			nums = nums[:i+1]
		}
	}
	return len(nums)
}
