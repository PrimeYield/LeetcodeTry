//242. Valid Anagram
package hashtable

func isAnagram(s string, t string) bool {
	countS := make(map[string]int)
	countT := make(map[string]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		countS[string(s[i])]++
		countT[string(t[i])]++
	}
	for str, i := range countS {
		if countT[str] != i {
			return false
		}
	}

	return true
}
