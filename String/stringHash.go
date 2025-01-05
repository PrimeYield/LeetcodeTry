//3271. Hash Divided String

package string

func stringHash(s string, k int) string {
	var result string
	alphabet := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
		"n": 13,
		"o": 14,
		"p": 15,
		"q": 16,
		"r": 17,
		"s": 18,
		"t": 19,
		"u": 20,
		"v": 21,
		"w": 22,
		"x": 23,
		"y": 24,
		"z": 25,
	}
	reAlphabet := map[int]string{
		0:  "a",
		1:  "b",
		2:  "c",
		3:  "d",
		4:  "e",
		5:  "f",
		6:  "g",
		7:  "h",
		8:  "i",
		9:  "j",
		10: "k",
		11: "l",
		12: "m",
		13: "n",
		14: "o",
		15: "p",
		16: "q",
		17: "r",
		18: "s",
		19: "t",
		20: "u",
		21: "v",
		22: "w",
		23: "x",
		24: "y",
		25: "z",
	}
	edge := len(s) / k
	sums := make([]int, 0, edge)
	for i := 0; i < edge; i++ {
		var sum int
		for j := 0; j < k; j++ {
			sum += alphabet[string(s[j])]
			if j == k-1 {
				sums = append(sums, sum)
				s = s[k:]
			}
		}
	}
	for i := 0; i < edge; i++ {
		result += reAlphabet[sums[i]%26]
	}
	return result
}

//Runtime 13ms Beats 9.68%
//Memory 9.02MB Beats 9.68%

// func stringHash(s string, k int) string {
// 	var intByS []int
// 	var sum, count int
// 	var sums = make([]int, 0, len(s)/k)
// 	for _, runeByS := range s {
// 		intByS = append(intByS, int(runeByS-'a'))
// 	}
// 	for {
// 		if len(intByS) == 0 {
// 			break
// 		}
// 		if count == k {
// 			sums = append(sums, sum)
// 			count = 0
// 			sum = 0
// 			continue
// 		}
// 		sum += intByS[0]
// 		intByS = intByS[1:]
// 		count++
// 	}
// 	var result string
// 	for i := 0; i < len(sums); i++ {
// 		trans := rune(sums[i] % 26)
// 		result += string(trans + 'a')
// 	}
// 	return result
// }
