//3304. Find the K-th Character in String Game I

package string

func kthCharacter(k int) byte {
	word := "a"
	for i := 0; i < k; i++ {
		if len(word) > k {
			return word[k-1]
		}
		temp := []byte(word)
		for j := 0; j < len(temp); j++ {
			temp[j] = temp[j] + 1
		}
		word += string(temp)
	}
	return word[k-1]
}

//0ms
//Memory 4.60MB
