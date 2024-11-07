// 169. Majority Element
package array

func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	var res int
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
	}
	for j, num := range countMap {
		if num > len(nums)/2 {
			res = j
		}
	}
	return res
}

/*
func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	var max, res int
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
	}
	for j, num := range countMap {
		if num > max {
			max = num
			res = j
		}

	}
	return res
}
*/
