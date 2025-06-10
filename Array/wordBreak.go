//139. Word Break

package array

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	// str := []byte(s)
	for i, j := 0, 0; j < len(s); j++ {
		if ok, _ := wordMap[s[i:j+1]]; ok {
			i = j + 1
			if i >= len(s) {
				return true
			}
		}
	}
	for i, j := len(s)-1, len(s)-1; j >= 0; j-- {
		if ok, _ := wordMap[s[j:i+1]]; ok {
			i = j - 1
			if i <= 0 {
				return true
			}
		}
	}
	return false
}

//要想辦法處理"aaaa","aa"這類型的
//s[j:i+1] == "aa"時會跳true
//永遠無法檢查"aaaa"的問題
