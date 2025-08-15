package array

import (
	"fmt"
	"strconv"
)

func largestGoodInteger(num string) string {
	res := ""
	max := 0
	for i := 0; i < len(num); {
		// if num[i] != num[i+1] {
		// 	i++
		// 	continue
		// }
		if i+2 < len(num) && num[i] == num[i+1] && num[i] == num[i+2] {
			temp, _ := strconv.Atoi(num[i : i+3])
			if temp > max {
				max = temp
			}
			res = fmt.Sprint(max)
			if res == "0" {
				res = "000"
			}
			i = i + 2
		} else {
			i++
		}
	}
	return res
}

//0ms
//4.17MB
