//27. Remove Element
package array

func removeElement(nums []int, val int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

/*
func removeElement(nums []int, val int) int {
	count := len(nums)
	for i, v := range nums {
		if v == val {
			nums[i] = 101
			count--
		}
	}

	sort.Ints(nums)
	return count
}
*/
