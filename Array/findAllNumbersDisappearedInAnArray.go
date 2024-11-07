//448. Find All Numbers Disappeared in an Array

package array

func findDisappearedNumbers(nums []int) []int {
	var res []int
	var allNum = make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		allNum[nums[i]] = nums[i]
	}
	allNum = allNum[1:]
	for i := 0; i < len(allNum); i++ {
		if allNum[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}
