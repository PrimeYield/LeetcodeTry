//1768. Merge Strings Alternately

package string

func mergeAlternately(word1 string, word2 string) string {
	end := min(len(word1), len(word2))
	res := ""
	for i := 0; i < end; i++ {
		res += string(word1[i]) + string(word2[i])
	}
	if len(word1) > end {
		res += word1[end:]
	} else if len(word2) > end {
		res += word2[end:]
	}
	return res
}

//2ms 59.59%
//4.19MB
