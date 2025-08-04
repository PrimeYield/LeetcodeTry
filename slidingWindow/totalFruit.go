//904. Fruit Into Baskets

package slidingwindow

func totalFruit(fruits []int) int {
	start := 0
	res := 0
	firstFruit := fruits[0]
	secondFruit := 0
	isInto := make([]bool, len(fruits))
	for i := 0; i < len(fruits)-1 && !isInto[i]; {
		firstFruit = fruits[i]
		currStart := start //固定起點
		isDiff := false
		for j := i + 1; j < len(fruits); j++ {
			// fmt.Println("i,",i,"j,",j,"fruits[j],",fruits[j],"start,",start)
			if fruits[j] != firstFruit && !isDiff {
				isDiff = true
				secondFruit = fruits[j]
				start = j //下個組合可能的起點
			} else if (fruits[j] != firstFruit && isDiff) && (fruits[j] != secondFruit && isDiff) {
				// fmt.Println("here")
				i = start
				length := j - currStart
				if res < length {
					res = length
				}
				break
			} else if j == len(fruits)-1 {
				length := j - currStart + 1
				if res < length {
					res = length
				}
			} else if fruits[j] == fruits[j-1] || fruits[j] == firstFruit || fruits[j] == secondFruit {
				continue
			}
		}
		isInto[currStart] = true //當過起點
	}
	return res
}

//9ms 93.56%
//9.72MB 62.66%
