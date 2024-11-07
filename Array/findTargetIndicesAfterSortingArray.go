//2089. Find Target Indices After Sorting Array

package array

func targetIndices(nums []int, target int) []int {
	var res []int
	for i, j := 0, len(nums)-1; i < j; j-- {
		if nums[i] > nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if j-i <= 1 {
			i++
			j = len(nums)
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res = append(res, i)
		}
	}
	return res
}
