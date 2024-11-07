//1920. Build Array from Permutation

package array

func buildArray(nums []int) []int {
	var res []int
	res = make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = nums[nums[i]]
	}
	return res
}
