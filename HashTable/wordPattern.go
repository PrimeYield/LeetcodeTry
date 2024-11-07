// 290. Word Pattern
package hashtable

func wordPattern(pattern string, s string) bool {
	s = s + " "
	sByte := []byte(s) //先將s每個單詞用切片存
	var sStringSlice []string
	for i := 0; i < len(sByte); {
		if sByte[i] == ' ' {
			sStringSlice = append(sStringSlice, string(sByte[:i]))
			sByte = sByte[i+1:]
			i = 0
		} else {
			i++
		}
	}
	if len(sStringSlice) != len(pattern) {
		return false
	}

	patternToSS := make(map[byte]string) //patern對應s的切片
	for i, _ := range pattern {
		if _, ok := patternToSS[pattern[i]]; ok {
			if sStringSlice[i] != patternToSS[pattern[i]] {
				return false
			}
		} else {
			patternToSS[pattern[i]] = sStringSlice[i]
		}
	}

	count := make(map[string]int)
	for _, sStr := range patternToSS {
		count[sStr]++
		if count[sStr] > 1 {
			return false
		}
	}
	return true
}
