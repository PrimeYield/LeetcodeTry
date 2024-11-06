package math

import "fmt"

func reverse(x int) int {
	k := 0
	j := 1
	var revSlice []int
	var res int
	for x > 0 {
		revSlice = append(revSlice, x%10)
		//k++
		x = x / 10
	}

	fmt.Println(revSlice)
	for i := 0; i < len(revSlice); i++ {
		for k = i; k < len(revSlice)-1; k++ {
			j = j * 10
		}
		res = res + revSlice[i]*j
		j = 1
	}
	if x < 0 {
		x = -x
		for x > 0 {
			revSlice = append(revSlice, x%10)
			//k++
			x = x / 10
		}

		for i := 0; i < len(revSlice); i++ {
			for k = i; k < len(revSlice)-1; k++ {
				j = j * 10
			}
			res = res + revSlice[i]*j
			j = 1
		}
		if res > 2147483648 {
			return 0
		}
		return -res
	}
	if res > 2147483648 {
		return 0
	}
	return res
}
