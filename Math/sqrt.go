// 69. Sqrt(x)
package math

func mySqrt(x int) int {
	var res, k int
	for i := 1; ; i++ {
		res = i * i
		if res > x {
			k = i - 1
			break
		}
	}
	return k
}
