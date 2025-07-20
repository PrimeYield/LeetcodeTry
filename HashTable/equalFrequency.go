//2423. Remove Letter To Equalize Frequency

package hashtable

func equalFrequency(word string) bool {
	count := make(map[string]int)
	for i := 0; i < len(word); i++ {
		count[string(word[i])]++
	}
	if len(count) == 1 { //word是由同一個字母組成
		return true
	}

	eleCount := make(map[int]int)
	for _, num := range count {
		eleCount[num]++
	}
	if len(eleCount) == 1 {
		for num := range eleCount {
			return num == 1
		}
	}
	if len(eleCount) > 2 {
		return false
	}
	if len(eleCount) == 2 {
		for _, num := range eleCount {
			if num == 1 {
				return true
			}
		}
		m := 0
		for num := range eleCount {
			m = max(num, m)
		}
		if _, exist := eleCount[m-1]; exist && eleCount[m] == 1 {
			return true
		}
	}
	return false
}

//0ms
//4MB
