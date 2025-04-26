// 62. Unique Paths
package dynamicprogramming

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				matrix[i] = append(matrix[i], 1)
				continue
			}
			if j == 0 {
				matrix[i] = append(matrix[i], 1)
				continue
			}
			sum := matrix[i-1][j] + matrix[i][j-1]
			matrix[i] = append(matrix[i], sum)
		}
	}
	return matrix[m-1][n-1]
}

//Runtime 0ms Beats 100%
//Memory 4.09MB Beats 55.74%
