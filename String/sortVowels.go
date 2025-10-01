//2785. Sort Vowels in a String

package string

import "sort"

func sortVowels(s string) string {
	//A > E > ... > U > a > e > ... > u
	temp := []byte(s)
	tempSlice := make([]byte, 0)
	tempNum := make([]int, 0)
	for i := 0; i < len(temp); i++ {
		if temp[i] == 'A' || temp[i] == 'E' || temp[i] == 'I' || temp[i] == 'O' || temp[i] == 'U' ||
			temp[i] == 'a' || temp[i] == 'e' || temp[i] == 'i' || temp[i] == 'o' || temp[i] == 'u' {
			tempSlice = append(tempSlice, temp[i])
			tempNum = append(tempNum, i)
		}
	}
	if len(tempSlice) == 0 {
		return s
	}
	sort.SliceStable(tempSlice, func(i, j int) bool {
		return tempSlice[i] < tempSlice[j]
	})
	for i := 0; i < len(temp); i++ {
		if len(tempNum) > 0 && i == tempNum[0] {
			temp[i] = tempSlice[0]
			tempNum = tempNum[1:]
			tempSlice = tempSlice[1:]
		}
	}
	return string(temp)
}
