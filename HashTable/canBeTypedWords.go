//1935. Maximum Number of Words You Can Type

package hashtable

func canBeTypedWords(text string, brokenLetters string) int {
	borkenMap := make(map[byte]bool)
	for i := 0; i < len(brokenLetters); i++ {
		borkenMap[brokenLetters[i]] = true
	}
	res := 0
	canType := true
	for i := 0; i < len(text); i++ {
		if text[i] != ' ' {
			if i == len(text)-1 {
				if borkenMap[text[i]] || !canType {
					break
				}
				res++
				break
			}
			if borkenMap[text[i]] {
				canType = false
				continue
			}
		}
		if text[i] == ' ' || i == len(text)-1 {
			if !canType {
				canType = true
				continue
			}
			res++
		}
	}
	return res
}
