//3255. Find the Power of K-Size Subarrays II

package array

/*能跑就是超時
func resultsArray(nums []int, k int) []int {
	var res = make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums)-k+1; i++ {
		cur := nums[i : i+k]
		res = append(res, isPower(cur, k))
	}
	return res
}
func isPower(nums []int, k int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			return -1
		}
	}
	return nums[k-1]
}
*/
func resultsArray(nums []int, k int) []int {
	if len(nums) < 2 || k == 1 {
		return nums
	}
	var res []int
	powerLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			powerLen++
		} else {
			powerLen = 1
		}
		if i >= k-1 {
			if powerLen >= k {
				res = append(res, nums[i])
			} else {
				res = append(res, -1)
			}
		}
	}
	return res
}
