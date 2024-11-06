package array

/*
func majorityElement(nums []int) []int {
	countMap := make(map[int]int)
	var res []int
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
	}
	for j, num := range countMap {
		if num > len(nums)/3 {
			res = append(res, j)
		}
	}
	return res
}
*/

func majorityElement(nums []int) []int {
	countMap := make(map[int]int)
	var res = make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
	}
	if len(countMap) >= (len(nums)-len(nums)/3)+1 {
		return res
	}
	if len(countMap) <= (len(nums)-(len(nums)/3)*2)+1 {
		for i, num := range countMap {
			if num > len(nums)/3 {
				res = append(res, i)
			}
		}
	}
	if len(countMap) <= (len(nums) - len(nums)/3) {
		for i, num := range countMap {
			if num > len(nums)/3 {
				res = append(res, i)
				return res
			}
		}
	}
	return res
}
