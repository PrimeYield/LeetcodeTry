//1394. Find Lucky Integer in an Array

package hashtable

func findLucky(arr []int) int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}
	luckNum := -1
	for target, count := range counts {
		if target == count {
			luckNum = max(luckNum, target)
		}
	}
	return luckNum
}

//0ms
//Memory 4.82MB
