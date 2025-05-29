//120. Triangle

package dynamicprogramming

import "math"

func minimumTotal(triangle [][]int) int {
	k := len(triangle) - 1
	for i := 1; i < len(triangle); i++ {
		// fmt.Println(triangle[i-1])
		if i == 1 {
			triangle[i][0] += triangle[0][0]
			triangle[i][1] += triangle[0][0]
			continue
		}
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][0]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				temp := min(triangle[i-1][j], triangle[i-1][j-1])
				triangle[i][j] += temp
			}
		}
	}
	res := math.MaxInt
	for i := 0; i < len(triangle[k]); i++ {
		res = min(res, triangle[k][i])
	}
	return res
}

//runtime 0ms Memory 4.82MB
