//166. Fraction to Recurring Decimal

package math

import "fmt"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator < 0 && denominator < 0 {
		//不需要特別處理，但要與只有一個負數分開處理
		numerator = -numerator
		denominator = -denominator
	}
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		if numerator < 0 {
			numerator = -numerator
		} else {
			denominator = -denominator
		}
		if numerator%denominator == 0 {
			return "-" + fmt.Sprint(numerator/denominator)
		}
		return "-" + fmt.Sprint(numerator/denominator) + "." + remainder((numerator%denominator), denominator)
	}
	if numerator%denominator == 0 {
		return fmt.Sprint(numerator / denominator)
	}
	return fmt.Sprint(numerator/denominator) + "." + remainder((numerator%denominator), denominator)
}

func remainder(num int, denominator int) string {
	res := ""
	// temp := num
	remainderMap := make(map[int]int)
	for i := 0; num > 0; i++ {
		// if i > 0 && num == temp {
		// 	res = "(" + res + ")"
		// 	return res
		// }
		if _, exist := remainderMap[num]; exist {
			tempHead := res[:remainderMap[num]]
			tempTail := res[remainderMap[num]:]
			res = tempHead + "(" + tempTail + ")"
			return res
		}
		remainderMap[num] = i
		num *= 10
		res += fmt.Sprint(num / denominator)
		num = num % denominator
	}
	return res
}
