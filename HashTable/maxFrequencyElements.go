//3005. Count Elements With Maximum Frequency

package hashtable

func maxFrequencyElements(nums []int) int {
	countMap := make(map[int]int)
	// countCount := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
	}
	// for _, count := range countMap {
	// 	countCount[count]++
	// }
	countSort := make([]int, 1)
	for _, count := range countMap {
		if countSort[0] < count {
			countSort = countSort[:1]
			countSort[0] = count
			// countSort = append(countSort, count)
			continue
		}
		if countSort[0] > count {
			continue
		}
		if countSort[0] == count {
			countSort = append(countSort, count)
		}
	}
	res := 0
	for i := 0; i < len(countSort); i++ {
		res += countSort[i]
	}
	return res
}
