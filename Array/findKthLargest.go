// 215. Kth Largest Element in an Array
package array

/*
func findKthLargest(nums []int, k int) int {
	var res []int
	res = make([]int, k, k)
	for i := 0; i < len(nums); i++ {
		for j := k - 1; j > 0; {
			if res[j] > nums[i] {
				j--
			} else if res[j] < nums[i] {
				var step []int
				step = res[1 : j+1]
				step = append(step, nums[i])
				step = append(step, res[j+1:]...)
				res = step
				break
			}
		}
	}
	return res[0]
}
*/
func findKthLargest(nums []int, k int) int {
	var max, min int = nums[0], nums[0] //避免最小值為0時不是nums內的元素
	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
		if min > nums[i] {
			min = nums[i]
		}
	}
	stay := make([]int, max-min+1) //用於計數int出現的次數
	for i := 0; i < len(nums); i++ {
		stay[nums[i]-min]++
	}
	for i := len(stay) - 1; i >= 0; i-- {
		k -= stay[i]
		if k <= 0 {
			return i + min
		}
	}
	return -1
}
