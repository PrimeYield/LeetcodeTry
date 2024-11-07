// 55. Jump Game
package array

import "fmt"

func canJump(nums []int) bool {
	length := len(nums)
	if nums[0] == 0 {
		if nums[0] == len(nums)-1 {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			fmt.Println(nums[:i+1])
			if check(nums[:i+1]) {
				fmt.Println(check(nums[:i+1]))
				length--
				if length <= 0 {
					return true
				}
				continue
			} else {
				fmt.Println(check(nums[:i+1]))
				return false
			}
		} else {
			length--
			if length <= 0 {
				return true
			}
		}

	}
	return false
}
func check(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > len(nums)-1 {
			return true
		}
	}
	return false
}
