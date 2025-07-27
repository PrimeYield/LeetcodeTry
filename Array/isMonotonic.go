//896. Monotonic Array

package array

// func isMonotonic(nums []int) bool {
// 	temp := make([]int, 0)
// 	for i := 0; i < len(nums); i++ {
// 		if i+1 < len(nums) && nums[i] == nums[i+1] {
// 			continue
// 		}
// 		temp = append(temp, nums[i])
// 	}
// 	for i := 2; i < len(temp); i++ {
// 		if temp[0] < temp[1] {
// 			if i+1 < len(temp) && temp[i] > temp[i+1] {
// 				return false
// 			}
// 		} else if temp[0] > temp[1] {
// 			if i+1 < len(temp) && temp[i] < temp[i+1] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
// //8ms Beats 3.80%
// //16.04MB
func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums); {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
			if i == len(nums)-1 {
				return true
			}
			continue
		}
		if i+1 < len(nums) && nums[i] < nums[i+1] {
			for j := i; j < len(nums); j++ {
				if j+1 < len(nums) && nums[j] > nums[j+1] {
					return false
				}
			}
			return true
		} else if i+1 < len(nums) && nums[i] > nums[i+1] {
			for j := i; j < len(nums); j++ {
				if j+1 < len(nums) && nums[j] < nums[j+1] {
					return false
				}
			}
			return true
		}
	}
	return true
}

//0ms
//10.74MB
