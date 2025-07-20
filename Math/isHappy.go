//202. Happy Number

package math

func isHappy(n int) bool {
	// for n > 0 {
	// 	curr += (n % 10) * (n % 10)
	// 	n /= 10
	// }
	retry := make(map[int]struct{})
	var loop func(n int) bool
	loop = func(n int) bool {
		if n == 1 {
			return true
		}
		if _, ok := retry[n]; ok {
			return false
		}
		retry[n] = struct{}{}
		curr := 0
		for n > 0 {
			curr += (n % 10) * (n % 10)
			n /= 10
		}
		return loop(curr)
	}
	return loop(n)
}

//0ms
//3.96MB 87.28%
