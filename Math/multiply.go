//43. Multiply Strings

package math

// func multiply(num1 string, num2 string) string {
// 	NUM1 := new(big.Int)
// 	NUM2 := new(big.Int)
// 	sum := new(big.Int)
// 	num1Int, _ := NUM1.SetString(num1, 10)
// 	num2Int, _ := NUM2.SetString(num2, 10)
// 	sum.Mul(num1Int, num2Int)
// 	return sum.String()
// }
// Runtime 0ms
// Memory 4.40MB
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]byte, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			store := res[i+j+1] + n1*n2
			res[i+j+1] = store % 10
			res[i+j] += store / 10
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}
	for i := 0; i < len(res); i++ {
		res[i] = '0' + res[i]
	}
	return string(res)
}

//Runtime 0ms Beats 100.00%
//Memory 4.22MB Beats 86.67%
