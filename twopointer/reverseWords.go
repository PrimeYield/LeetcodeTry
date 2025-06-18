//151. Reverse Words in a String

package twopointer

func reverseWords(s string) string {
	stack := make([]string, 0)
	for i, j := 0, 0; j < len(s); j++ {
		if s[j] == ' ' {
			if s[i] == ' ' {
				i = j + 1
				continue
			}
			stack = append(stack, s[i:j])
			i = j + 1
		}
		if j == len(s)-1 && s[j] != ' ' {
			stack = append(stack, s[i:])
		}
	}
	res := ""
	for len(stack) > 0 {
		if len(stack) > 1 {
			res += stack[len(stack)-1] + " "
		} else {
			res += stack[len(stack)-1]
		}
		stack = stack[:len(stack)-1]
	}
	return res
}

//runtime 3ms
//Memory 8.89MB
