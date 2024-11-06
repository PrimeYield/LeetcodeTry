package array

func missingNumber(nums []int) int {
	var sum, k, arraySum int = 0, 1, 0
	for i := 0; i < len(nums); i++ {
		arraySum = arraySum + nums[i]
		sum = sum + k
		k++
	}
	return sum - arraySum
}
