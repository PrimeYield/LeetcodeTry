// 217. Contains Duplicate
package array

func containsDuplicate(nums []int) bool {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
		if countMap[num] == 2 {
			return true
		}
	}

	return false
}
