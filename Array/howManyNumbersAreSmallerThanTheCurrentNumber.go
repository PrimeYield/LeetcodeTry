//1365. How Many Numbers Are Smaller Than the Current Number

package array

func smallerNumbersThanCurrent(nums []int) []int {
	valueMap := make(map[int]int)
	var res []int
	for i := range nums {
		valueMap[i] = nums[i]
	}
	for _, num := range nums {
		count := 0
		for _, num2 := range valueMap {
			if num2 < num {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}
