//2390. Removing Stars From a String

package stack

func removeStars(s string) string {
	stack := make([]rune, 0)
	for _, str := range s {
		if str == '*' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, str)
	}
	res := ""
	for i := 0; i < len(stack); i++ {
		res += string(stack[i])
	}
	return res
}

//runtime 2094ms
//Memory 97.4MB
