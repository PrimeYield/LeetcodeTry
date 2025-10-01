//2197. Replace Non-Coprime Numbers in Array

package math

func replaceNonCoprimes(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		for len(res) > 0 {
			last := res[len(res)-1]
			if gcd(last, num) > 1 {
				num = lcm(last, num)
				res = res[:len(res)-1]
			} else {
				break
			}
		}
		res = append(res, num)
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return (a * b) / gcd(a, b)
}
