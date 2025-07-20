//258. Add Digits

package math

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	temp := 0
	for num > 0 {
		temp += num % 10
		num /= 10
	}
	return addDigits(temp)
}

//0ms
//4.06MB 64.44%
