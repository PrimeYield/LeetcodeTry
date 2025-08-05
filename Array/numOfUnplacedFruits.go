//3477. Fruits Into Baskets II

package array

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	isUsed := make([]bool, len(baskets))
	for i := 0; i < len(fruits); i++ {
		for j := 0; j < len(baskets); j++ {
			if baskets[j] >= fruits[i] && !isUsed[j] {
				isUsed[j] = true
				break
			}
		}
	}
	res := 0
	for i := 0; i < len(isUsed); i++ {
		if !isUsed[i] {
			res++
		}
	}
	return res
}

//0ms
//6.01MB 55.88%
