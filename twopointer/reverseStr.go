//541. Reverse String II

package twopointer

func reverseStr(s string, k int) string {
	count := 0
	res := ""
	for i, j := 0, k; j < len(s); i, j = i+k, j+k {
		if count%2 == 0 {
			//as -> sa
			temp := []byte(s[i:j])
			str := ""
			for m := 0; m < len(temp); m++ {
				str = string(temp[m]) + str
			}
			res += str
		} else {
			res += string(s[i:j])
		}
		count++
	}
	if len(res) < len(s) {
		if count%2 == 0 {
			temp := []byte(s[len(res):])
			for i := len(temp) - 1; i >= 0; i-- {
				res += string(temp[i])
			}
		} else {
			res += s[len(res):]
		}
	}
	return res
}

//runtime 7ms Beats 10.57%
//Memory 9.53MB Beats 5.69%
