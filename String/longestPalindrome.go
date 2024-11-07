// 5. Longest Palindromic Substring
package string

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var allPalindrome []string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if boolPalindrome(s[i : j+1]) {
				allPalindrome = append(allPalindrome, s[i:j+1])
			}
		}
	}
	if len(allPalindrome) == 0 {
		return string(s[0])
	}
	var m, longest int
	for k := 0; k < len(allPalindrome); k++ {
		if longest < len(allPalindrome[k]) {
			longest = len(allPalindrome[k])
			m = k
		}
	}
	return allPalindrome[m]
}
func boolPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		k := 0
		if s[i] == s[j] {
			k++
			i = i + k
			j = j - k
		} else {
			return false
		}
	}
	return true
}
