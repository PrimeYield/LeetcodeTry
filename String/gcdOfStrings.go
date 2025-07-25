//1071. Greatest Common Divisor of Strings

package string

func gcdOfStrings(str1 string, str2 string) string {
	gcdLength := 0
	len1 := len(str1)
	len2 := len(str2)
	temp := ""
	if len1 > len2 {
		gcdLength = len2
	} else {
		gcdLength = len1
	}
	for ; gcdLength > 0; gcdLength-- {
		check1 := true
		check2 := true
		if len1%gcdLength == 0 && len2%gcdLength == 0 {
			temp = str2[:gcdLength]
			for i := 0; i < len1; i = i + gcdLength {
				if str1[i:i+gcdLength] != temp {
					check1 = false
					break
				}
			}
			for i := 0; i < len2; i = i + gcdLength {
				if str2[i:i+gcdLength] != temp {
					check2 = false
					break
				}
			}
			if !check1 || !check2 {
				temp = ""
			}
			if len(temp) > 0 {
				return temp
			}
		}
	}
	return ""
}

//0ms Beats 100%
//4.25MB Beats 95.67%
