// 13. Roman to Integer
package hashtable

import "fmt"

func romanToInt(s string) int {
	trans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var sum int

	for i := len(s) - 1; i >= 0; i-- {

		sum = sum + trans[string(s[i])]
		if i > 0 && trans[string(s[i])] > trans[string(s[i-1])] {
			sum = sum - 2*trans[string(s[i-1])]
		}
	}
	fmt.Print(sum)
	return sum
}
