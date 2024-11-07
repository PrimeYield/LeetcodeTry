package string

func scoreOfString(s string) int {
	var sum int
	for i := 0; i < len(s)-1; i++ {
		cur := int(s[i] - s[i+1])
		if cur < 0 {
			sum = sum - cur
			continue
		}
		sum = sum + cur
	}
	return sum
}
