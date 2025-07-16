//3136. Valid Word

package string

func isValid(word string) bool {
	//'0' = 48
	//'9' = 57
	//'a' = 97	'A' =65
	//'e' = 101 'E' =69
	//'i' = 105 'I' =73
	//'o' = 111 'O' =79
	//'u' = 117 'U' = 85
	//'z' = 122 'Z' = 90
	if len(word) < 3 {
		return false
	}
	vowel := 0
	consonant := 0
	for i := 0; i < len(word); i++ {
		if (word[i] >= 97 && word[i] <= 122) || (word[i] >= 65 && word[i] <= 90) {
			if word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] == 'u' || word[i] == 'A' || word[i] == 'E' || word[i] == 'I' || word[i] == 'O' || word[i] == 'U' {
				vowel++
			} else {
				consonant++
			}
		} else if word[i] >= 48 && word[i] <= 57 {
			continue
		} else {
			return false
		}
	}
	if vowel < 1 || consonant < 1 {
		return false
	}
	return true
}

//0ms 4.01MB
