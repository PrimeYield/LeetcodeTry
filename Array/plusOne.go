//66. Plus One
package array

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		digits[i]++
		if digits[i]/10 > 0 {
			digits[i] = 0
			if i-1 >= 0 {
				continue
			} else if i == 0 {
				indexZero := []int{1}
				digits = append(indexZero, digits...)
			}
		} else {
			break
		}
	}
	return digits
}
