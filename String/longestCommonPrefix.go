package string

func longestCommonPrefix(strs []string) string {
	var res string
	var runMin, strStart int
	runMin = 999
	for i := 0; i < len(strs); i++ {
		if runMin > len(strs[i]) {
			runMin = len(strs[i])
			strStart = i
		}
	}
	res = strs[strStart] //this is a stack
	for i := 0; i < len(strs); i++ {
		for j := runMin - 1; j >= 0; j-- {
			if strs[i][j] != res[j] {
				res = res[:j]
				runMin = len(res)
			}
		}
	}
	return res
}
