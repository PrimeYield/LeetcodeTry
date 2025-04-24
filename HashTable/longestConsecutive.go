//128. Longest Consecutive Sequence

package hashtable

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}
	longest := 0
	for num, _ := range numsMap {
		if !numsMap[num-1] {
			curr := num
			currLen := 1
			for numsMap[curr+1] {
				curr++
				currLen++
			}
			if longest < currLen {
				longest = currLen
			}
		}
	}
	return longest
}
