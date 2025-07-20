//1945. Sum of Digits of String After Convert

package string

import "fmt"

var trans map[string]string

func getLucky(s string, k int) int {
	trans = map[string]string{
		"a": "1", "b": "2", "c": "3", "d": "4", "e": "5",
		"f": "6", "g": "7", "h": "8", "i": "9", "j": "10",
		"k": "11", "l": "12", "m": "13", "n": "14", "o": "15",
		"p": "16", "q": "17", "r": "18", "s": "19", "t": "20",
		"u": "21", "v": "22", "w": "23", "x": "24", "y": "25", "z": "26",
	}
	transStr := ""
	for i := 0; i < len(s); i++ {
		transStr += trans[string(s[i])]
	}
	ans := 0
	byteStr := []byte(transStr)
	for i := 0; i < k; i++ {
		res := 0
		for j := 0; j < len(byteStr); j++ {
			res = res + int(byteStr[j]) - 48
		}
		// res += 48
		transStr = fmt.Sprint(res)
		byteStr = []byte(transStr)
		ans = res
	}
	return ans
}

//0ms
//5.15MB 26.09%
