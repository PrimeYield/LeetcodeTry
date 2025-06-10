// 400. Nth Digit
package math

func findNthDigit(n int) int {
	startNum := 1
	digit := 1
	count := 9
	n = n - 1
	for {
		if n < count*digit {
			targetNum := startNum + n/digit
			for i := 1; i < digit-(n%digit); i++ {
				targetNum /= 10
			}
			return targetNum % 10
		}
		n -= digit * count
		startNum *= 10
		digit++
		count *= 10
	}
}

//runtime 0ms
//Memory 3.97MB
