// 1304. Find N Unique Integers Sum up to Zero

package math

func sumZero(n int) []int {
	res := []int{}
	if n%2 == 1 {
		for i := 0; i <= (n / 2); i++ {
			res = append(res, i)
			if i > 0 {
				res = append(res, -i)
			}
		}
		return res
	}
	if n%2 == 0 {
		for i := 1; i <= (n / 2); i++ {
			res = append(res, i)
			res = append(res, -i)
		}
		return res
	}
	return res
}
