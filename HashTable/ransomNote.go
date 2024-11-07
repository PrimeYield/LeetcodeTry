// 383. Ransom Note
package hashtable

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	countMap := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		countMap[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if countMap[ransomNote[i]] == 0 {
			return false
		}
		countMap[ransomNote[i]]--
	}
	return true
}
