//1207. Unique Number of Occurrences

package hashtable

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	appTimes := make(map[int]int)
	for _, counts := range count {
		appTimes[counts]++
	}
	for _, times := range appTimes {
		if times > 1 {
			return false
		}
	}
	return true
}

//runtime 0ms
//Memory 4.4MB
