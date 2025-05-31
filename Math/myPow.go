//50. Pow(x, n)

package math

import "math"

/*
func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
     if x == -1 {
        if n %2 == 1 {
            return -1
        }
        return 1
    }
    if n < -2147483647 || n > 2147483646 {
        return 0
    }

	res := x
	if n < 0 {
		for i := 0; i <= n*(-1); i++ {
			res /= x
		}
	} else {
		for i := 1; i < n; i++ {
			res *= x
		}
	}
	return res
}
runtime 784ms beats2.41%
Memory 4.07MB beats 19.92%
*/

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

//runtime 0ms Memory 3.99MB
