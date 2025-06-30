//594. Longest Harmonious Subsequence

package hashtable

func findLHS(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	res := 0
	for num, count := range counts {
		if _, exist := counts[num+1]; exist {
			if res < counts[num]+counts[num+1] {
				res = count + counts[num+1]
			}
		}
	}
	return res
}

//runtime 18ms
//Memory 9.08MB
