//912. Sort an Array

package array

func sortArray(nums []int) []int {
	max := nums[0]
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		if min > nums[i] {
			min = nums[i]
		}
	}
	length := max - min + 1
	temp := make([]int, length)
	for i := 0; i < len(nums); i++ {
		temp[nums[i]-min]++
	}
	var res []int
	for i := 0; i < len(temp); i++ {
		if temp[i] > 0 {
			for j := 0; j < temp[i]; j++ {
				res = append(res, i+min)
			}
		}
	}
	return res
}

//runtime 13ms Beats 83.51%
//Memory 9.10MB Beats 91.10%
