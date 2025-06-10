//139. Word Break

package dynamicprogramming

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, str := range wordDict {
		wordMap[str] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

//runtime 0ms
//Memory 4.08MB Beats 85.97%
