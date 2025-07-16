//2734. Lexicographically Smallest String After Substring Operation

package string

func smallestString(s string) string {
	start := 0
	end := 0
	temp := []byte(s)
	for i := 0; i < len(s); i++ {
		for s[start] == 'a' {
			if start == len(s)-1 {
				temp[start] = 'z'
				return string(temp)
			}
			start++
		}
		if temp[i] == 'a' && i > start {
			end = i
			break
		} else {
			end = i + 1
		}
	}
	for i := start; i < end; i++ {
		temp[i] -= 1
	}
	return string(temp)
}

//2ms Beats 100%
//9.73MB  Beats 80.00%
