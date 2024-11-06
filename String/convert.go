package string

func convert(s string, numRows int) string {
	var res string
	firstRowNextIndex := (numRows - 1) * 2
	if len(s) < 3 || numRows == 1 {
		return s
	}
	for r := 1; r <= numRows; r++ {
		for i := 0; i < len(s); i++ {
			if firstRowNextIndex > 0 && i%firstRowNextIndex == r-1 {
				res = res + string(s[i])
			}
			if firstRowNextIndex > 0 && i%firstRowNextIndex > r-1 && (i+(r-1)*2)%firstRowNextIndex == r-1 {
				res = res + string(s[i])
			}

		}
	}
	return res
}
