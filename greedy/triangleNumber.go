//611. Valid Triangle Number

package greedy

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := len(nums) - 1; j < k; k-- {
				if nums[i]+nums[j] <= nums[k] {
					continue
				}
				if nums[i]+nums[j] > nums[k] {
					res++
				}
			}
		}
	}
	return res
}

//1651ms 4.95MB
// func triangleNumber(nums []int) int {
// 	sort.Ints(nums)
// 	res := 0
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			num := nums[i] + nums[j]
// 			temp := nums[j+1:]
// 			mid := len(temp) / 2
// 			k := binarySearch(temp, mid, num)
// 			res += k
// 		}
// 	}
// 	return res
// }

// func binarySearch(nums []int, mid, num int) int {
// 	if num < nums[len(nums)-1] {
// 		return len(nums) - 1
// 	}
// 	if num > nums[mid] {
// 		nums = nums[mid:]
// 		mid2 := len(nums) / 2
// 		return binarySearch(nums, mid2, num)
// 	}
// 	return mid
// }
