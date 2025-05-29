// 8. String to Integer (atoi)
package string

func myAtoi(s string) int {
	unit := []byte(s)
	k := 0
	for i := 0; i < len(unit); i++ {
		if unit[i] == ' ' {
			k++
		}
		if unit[i] != ' ' {
			unit = unit[k:]
			break
		}
	}
	res := 0
	sign := 1
	for i := 0; i < len(unit); i++ {
		if i == 0 && unit[i] == '-' {
			sign = -1
			continue
		} else if i == 0 && unit[i] == '+' {
			sign = 1
			continue
		}
		if unit[i] != '0' && unit[i] != '1' && unit[i] != '2' && unit[i] != '3' && unit[i] != '4' && unit[i] != '5' && unit[i] != '6' && unit[i] != '7' && unit[i] != '8' && unit[i] != '9' {
			break
		}
		res = res*10 + int(unit[i]-'0')
		if res*sign > 2147483647 {
			return 2147483647
		}
		if res*sign < -2147483648 {
			return -2147483648
		}
	}
	// fmt.Println(res)

	return res * sign
}

//runtime 0ms Memory 4.08 MB
