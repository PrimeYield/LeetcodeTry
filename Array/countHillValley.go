//2210. Count Hills and Valleys in an Array

package array

func countHillValley(nums []int) int {
	res := 0
	temp := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			continue
		}
		temp = append(temp, nums[i])
	}
	for i := 1; i < len(temp)-1; i++ {
		if temp[i-1] < temp[i] && temp[i] > temp[i+1] {
			res++
		} else if temp[i-1] > temp[i] && temp[i] < temp[i+1] {
			res++
		}
	}
	return res
}

//0ms
//4.26MB
