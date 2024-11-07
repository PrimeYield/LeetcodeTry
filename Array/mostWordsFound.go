//2114. Maximum Number of Words Found in Sentences

package array

func mostWordsFound(sentences []string) int {
	count := 0
	for i := 0; i < len(sentences); i++ {
		k := 0
		for j := 0; j < len(sentences[i]); j++ {
			if j == 0 {
				k = 0
			}
			if sentences[i][j] == ' ' {
				k++
			}
		}
		if count < k {
			count = k
		}
	}
	return count + 1
}
