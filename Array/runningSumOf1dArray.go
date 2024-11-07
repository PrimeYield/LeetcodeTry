//1480. Running Sum of 1d Array

package array

func runningSum(nums []int) []int {
	var res []int
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		res = append(res, sum)
	}
	return res
}
