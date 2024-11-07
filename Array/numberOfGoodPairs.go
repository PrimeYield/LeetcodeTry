//1512. Number of Good Pairs

package array

func numIdenticalPairs(nums []int) int {
	var count int
	for i, j := 0, len(nums)-1; i < j; j-- {
		if nums[i] == nums[j] {
			count++
		}
		if j-i <= 1 {
			j = len(nums)
			i++
		}
	}
	return count
}
