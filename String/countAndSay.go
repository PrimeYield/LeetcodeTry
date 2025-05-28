// 38. Count and Say
package string

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	var count [10]int
	var res string
	for i, unit := range str {
		j, _ := strconv.Atoi(string(unit))
		count[j]++
		// fmt.Println("array",count)
		if (i+1 < len(str) && str[i] != str[i+1]) || i == len(str)-1 {
			res += strconv.Itoa(count[j]) + strconv.Itoa(j)
			// fmt.Println("res",res)
			count[j] = 0
		}
	}
	return res
}

//runtime 8ms memory 11.66MB
