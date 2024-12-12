//2433. Find The Original Array of Prefix Xor

package array

// func findArray(pref []int) []int {
// 	var res []int
// 	var cur int
// 	res = append(res, pref[0])
// 	for i := 1; i < len(pref); i++ {
// 		cur = pref[i-1] ^ pref[i]
// 		res = append(res, cur)
// 	}
// 	return res
// }

func findArray(pref []int) []int {
	res := make([]int, len(pref), len(pref))
	var cur int
	res[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		cur = pref[i-1] ^ pref[i]
		res[i] = cur
	}
	return res
}
