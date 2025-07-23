//1717. Maximum Score From Removing Substrings

package stack

func maximumGain(s string, x int, y int) int {
	// temp := ""
	temp := make([]byte, 0)
	high := max(x, y)
	low := min(x, y)
	single := 0 //沒配對到的低分起頭
	res := 0
	for i := 0; i < len(s); i++ {
		if x > y {
			if s[i] == 'a' {
				temp = append(temp, s[i])
			}
		} else {
			if s[i] == 'b' {
				temp = append(temp, s[i])
			}
		}
		if s[i] != 'a' && s[i] != 'b' {
			if len(temp) > 0 && single > 0 {
				res += min(len(temp), single) * low
			}
			temp = temp[:0]
			single = 0
		} else if len(temp) > 0 && s[i] != temp[len(temp)-1] {
			res += high
			temp = temp[:len(temp)-1]
		} else if len(temp) == 0 {
			single++
		}
		if i == len(s)-1 && len(temp) > 0 && single > 0 {
			res += min(len(temp), single) * low
		}
	}
	return res
}

//14ms Beats 100%
//8.82MB Beats 80%
