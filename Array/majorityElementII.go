// 229. Majority Element II
package array

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
