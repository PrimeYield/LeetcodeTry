// 205. Isomorphic Strings
package hashtable

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sCorrespondsToT := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := sCorrespondsToT[s[i]]; ok {
			if t[i] != sCorrespondsToT[s[i]] {
				return false
			}
		} else {
			sCorrespondsToT[s[i]] = t[i]
		}

	}
	count := make(map[byte]int)
	for _, tByte := range sCorrespondsToT {

		count[tByte]++
		fmt.Println(tByte, count)
		if count[tByte] > 1 {
			return false
		}
	}
	return true
}
