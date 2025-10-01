//3541. Find Most Frequent Vowel and Consonant

package hashtable

func maxFreqSum(s string) int {
	countMap := make(map[byte]int)
	sByte := []byte(s)
	for i := 0; i < len(sByte); i++ {
		countMap[sByte[i]]++
	}
	maxVowel := 0
	maxConsonant := 0
	for letter, count := range countMap {
		if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
			maxVowel = max(maxVowel, count)
		} else {
			maxConsonant = max(maxConsonant, count)
		}
	}
	return maxVowel + maxConsonant
}
