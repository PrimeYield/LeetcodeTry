//1323. Maximum 69 Number

package greedy

import "fmt"

func maximum69Number(num int) int {
	str := fmt.Sprint(num)
	strByte := []byte(str)
	j := 0
	res := 0
	for i := 0; j < len(strByte) && i < len(strByte); i++ {
		if strByte[i] == '6' {
			strByte[i] = '9'
			for j = i; j < len(strByte); j++ {
				res = res*10 + int(strByte[j]-'0')
			}
			break
		}
		res = res*10 + int(strByte[i]-'0')
	}
	return res
}

//0ms
//3.87MB 75.71%
