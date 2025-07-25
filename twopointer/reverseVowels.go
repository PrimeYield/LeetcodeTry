// 345. Reverse Vowels of a String
package twopointer

func reverseVowels(s string) string {
	sByte := []byte(s)
	for i, j := 0, len(sByte)-1; i < j; {
		if sByte[i] != 'a' && sByte[i] != 'e' && sByte[i] != 'i' && sByte[i] != 'o' && sByte[i] != 'u' &&
			sByte[i] != 'A' && sByte[i] != 'E' && sByte[i] != 'I' && sByte[i] != 'O' && sByte[i] != 'U' {
			i++
		}
		if sByte[j] != 'a' && sByte[j] != 'e' && sByte[j] != 'i' && sByte[j] != 'o' && sByte[j] != 'u' &&
			sByte[j] != 'A' && sByte[j] != 'E' && sByte[j] != 'I' && sByte[j] != 'O' && sByte[j] != 'U' {
			j--
		}
		if (sByte[i] == 'a' || sByte[i] == 'e' || sByte[i] == 'i' || sByte[i] == 'o' || sByte[i] == 'u' || sByte[i] == 'A' || sByte[i] == 'E' || sByte[i] == 'I' || sByte[i] == 'O' || sByte[i] == 'U') &&
			(sByte[j] == 'a' || sByte[j] == 'e' || sByte[j] == 'i' || sByte[j] == 'o' || sByte[j] == 'u' || sByte[j] == 'A' || sByte[j] == 'E' || sByte[j] == 'I' || sByte[j] == 'O' || sByte[j] == 'U') {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			i++
			j--
		}
	}
	return string(sByte)
}

//0ms
//6.01MB Beats 83.95%
