package string

func firstUniqChar(s string) int {
	countMap := make(map[byte]int)
	var sorting []int
	for i, _ := range s {
		countMap[s[i]]++
	}
	for i, _ := range s {
		if countMap[s[i]] == 1 {
			sorting = append(sorting, i)
		}
	}
	if len(sorting) > 0 {
		return sorting[0]
	}
	return -1
}
