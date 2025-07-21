//1957. Delete Characters to Make Fancy String

package string

func makeFancyString(s string) string {
	// res := ""
	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(res) >= 2 && res[len(res)-2] == s[i] && res[len(res)-1] == s[i] {
			continue
		}
		res = append(res, s[i])
	}
	return string(res)
}

//17ms Beats 52.17%
//9.46MB Beats 65.22%
