package string

func lengthOfLastWord(s string) int {
	var str []string
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == 32 {
			s = s[:i]
		} else {
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			str = append(str, string(s[i]))
		} else {
			break
		}
	}
	return len(str)
}
