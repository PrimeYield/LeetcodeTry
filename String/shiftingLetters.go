//848. Shifting Letters

package string

func shiftingLetters(s string, shifts []int) string {
	temp := []byte(s)
	tempShifts := make([]int, len(s))
	tempSum := 0
	for i := 0; i < len(s); i++ {
		tempSum += shifts[i]
	}
	tempShifts[0] = tempSum
	for i := 1; i < len(s); i++ {
		tempShifts[i] = (tempShifts[i-1] - shifts[i-1])
	}
	for i := 0; i < len(s); i++ {
		tempShifts[i] %= 26
	}
	for i := 0; i < len(s); i++ {
		if tempShifts[i] == 0 {
			continue
		}
		temp[i] = (temp[i] + byte(tempShifts[i]))
		if temp[i] > 'z' {
			temp[i] -= 26
		}
	}
	return string(temp)
}

//runtime 7ms
//Memory 10.32MB
