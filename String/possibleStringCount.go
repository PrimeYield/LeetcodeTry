//3330. Find the Original Typed String I

package string

func possibleStringCount(word string) int {
	res := 1
	for i := 0; i < len(word); i++ {
		if i < len(word)-1 && word[i] == word[i+1] {
			res++
		}
	}
	return res
}

//runtime 0ms
//Memory 4.06MB
