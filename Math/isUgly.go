//263. Ugly Number

package math

func isUgly(n int) bool {
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			continue
		}
		if n%3 == 0 {
			n /= 3
			continue
		}
		if n%5 == 0 {
			n /= 5
			continue
		}
		if n%2 != 0 && n%3 != 0 && n%5 != 0 {
			return false
		}
	}
	return n == 1
}

//0ms
//4.07MB 67.74%
