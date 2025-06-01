//605. Can Place Flowers

package array

// func canPlaceFlowers(flowerbed []int, n int) bool {
// 	if len(flowerbed) == 1 {
// 		if flowerbed[0] == 0 {
// 			return n <= 1
// 		} else {
// 			return n == 0
// 		}
// 	}
// 	count := 0
// 	for i := 0; i < len(flowerbed); i++ {
// 		if flowerbed[i] == 1 {
// 			continue
// 		} else if i == 0 {
// 			if flowerbed[i] != 1 {
// 				if flowerbed[i+1] == 1 {
// 					continue
// 				} else {
// 					flowerbed[i] = 1
// 					count++
// 				}
// 			} else {
// 				continue
// 			}
// 		} else if i == len(flowerbed)-1 {
// 			if flowerbed[i] != 1 {
// 				if flowerbed[i-1] == 1 {
// 					break
// 				} else {
// 					flowerbed[i] = 1
// 					count++
// 					break
// 				}
// 			}
// 		} else {
// 			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
// 				flowerbed[i] = 1
// 				count++
// 			} else {
// 				continue
// 			}
// 		}
// 	}
// 	return count >= n
// }

//runtime 4ms beats 6.31%
//Memory 7.95MB Beats 79.13%
func canPlaceFlowers(flowerbed []int, n int) bool {
	lenght := len(flowerbed)
	for i := 0; i < lenght; i++ {
		if n == 0 {
			return true
		}
		if (i == 0 || flowerbed[i-1] == 0) && flowerbed[i] == 0 && (i == lenght-1 || flowerbed[i+1] == 0) {
			n--
			i++
			if n == 0 {
				return true
			}
		}
	}
	return n <= 0
}

//runtime 0ms
//Memory 8.26MB beats 36.35%
