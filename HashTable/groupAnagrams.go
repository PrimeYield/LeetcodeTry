// 49. Group Anagrams
package hashtable

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	classification := make(map[string][]string)
	var res [][]string
	for i := 0; i < len(strs); i++ {
		classification[sortLetter(strs[i])] = append(classification[sortLetter(strs[i])], strs[i])
	}
	for _, strss := range classification {
		res = append(res, strss)
	}
	return res
}

func sortLetter(str string) string {
	var res string
	strSlice := strToSlice(str)
	sort.Strings(strSlice)
	for i := 0; i < len(strSlice); i++ {
		res += strSlice[i]
	}
	return res
}

func strToSlice(str string) []string {
	var res []string
	for i := 0; i < len(str); i++ {
		res = append(res, string(str[i]))
	}
	return res
}

/*
Runtime 51ms Beats 5.95%
Memory 10.50MB Beats 29.25%
*/
