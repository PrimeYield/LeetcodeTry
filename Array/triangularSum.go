//2221. Find Triangular Sum of an Array

package array

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	temp := []int{}
	for i := 0; i < len(nums)-1; i++ {
		temp = append(temp, (nums[i]+nums[i+1])%10)
	}
	return triangularSum(temp)
}
