//334. Increasing Triplet Subsequence

package greedy

import "math"

// func increasingTriplet(nums []int) bool {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] > nums[j] {
// 				break
// 			}
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[i] < nums[j] && nums[j] < nums[k] {
// 					return true
// 				}
// 			}
// 		}
// 		if i == len(nums)-1 {
// 			return false
// 		}

// 	}
// 	return true
// } 超時
// nums[i] == 0 nums[j] == 0 造成 for i進入for j後馬上break
// for i 跟 for j在做無意義的輪詢

//要怎麼最快的確定前兩個數??
// func increasingTriplet(nums []int) bool {
// 	nums_i := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] > nums_i {
// 			temp := nums[i]
// 			for k := i + 1; k < len(nums); k++ {
// 				if temp < nums[k] {
// 					return true
// 				}
// 			}
// 		} else if nums[i] < nums_i {
// 			nums_i = nums[i]
// 		} else {
// 			continue
// 		}
// 	}
// 	return false
// }超時

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	nums_i := nums[0]
	nums_j := math.MaxInt32
	for i := 1; i < n; i++ {
		if nums[i] > nums_j {
			return true
		}
		if nums[i] > nums_i {
			nums_j = nums[i]
		} else if nums[i] < nums_i {
			nums_i = nums[i]
		}
	}
	return false
}

//0ms
//18.75MB
