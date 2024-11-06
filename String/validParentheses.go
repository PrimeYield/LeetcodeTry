package string

func close(stack []byte) bool {
	l := len(stack)
	if l <= 1 {
		return false
	}
	if stack[l-1] == ')' && stack[l-2] == '(' {
		return true
	}
	if stack[l-1] == ']' && stack[l-2] == '[' {
		return true
	}
	if stack[l-1] == '}' && stack[l-2] == '{' {
		return true
	}

	return false
}

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if s[i] == ')' && close(stack) || s[i] == ']' && close(stack) || s[i] == '}' && close(stack) {
			stack = stack[:len(stack)-2]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
