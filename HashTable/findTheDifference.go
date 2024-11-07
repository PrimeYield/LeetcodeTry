//389. Find the Difference

package hashtable

func findTheDifference(s string, t string) byte {
	countMapS := make(map[byte]int)
	countMapT := make(map[byte]int)
	var res byte
	for _, str := range s {
		countMapS[byte(str)]++
	}
	for _, str := range t {
		countMapT[byte(str)]++
	}
	for key, value := range countMapT {
		if value != countMapS[key] {
			res = key
		}
	}
	return res
}
