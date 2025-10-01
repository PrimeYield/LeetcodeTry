//1317. Convert Integer to the Sum of Two No-Zero Integers

package math

func getNoZeroIntegers(n int) []int {
	res := []int{}
	for i := 1; i < n; i++ {
		temp := n
		if i%10 == 0 {
			continue
		}
		if i%10 != 0 {
			dummy := i
			for dummy/10 > 0 {
				dummy /= 10
				if dummy%10 == 0 {
					break
				}
			}
			if dummy%10 == 0 {
				continue
			}
		}
		temp -= i
		for temp/10 > 0 {
			if temp%10 == 0 {
				break
			}
			temp /= 10
		}
		if temp < 10 {
			res = append(res, i)
			res = append(res, n-i)
			break
		}
	}
	return res
}
