//3. Longest Substring Without Repeating Characters
package hashtable

func lengthOfLongestSubstring(s string) int {
	countMap := make(map[rune]int)
	var resSlice, delSlice []rune
	l := 0
	maxLen := 0
	for _, letter := range s {
		countMap[letter]++
		resSlice = append(resSlice, letter)
		//fmt.Println(resSlice)
		if countMap[letter] > 1 {
			i := 0
			for _, reLetter := range resSlice {
				if reLetter != letter {
					i++
				} else if reLetter == letter {
					break
				}
			}
			delSlice = resSlice[:i+1]
			//fmt.Println(delSlice)
			for _, delLetter := range delSlice {
				countMap[delLetter]--
			}
			resSlice = resSlice[i+1:]
			//fmt.Println(resSlice)
		}
		l = len(resSlice)
		if l > maxLen {
			maxLen = l
		}

	}

	return maxLen
}
