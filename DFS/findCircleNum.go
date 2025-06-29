// 547. Number of Provinces
package dfs

func findCircleNum(isConnected [][]int) int {
	cities := len(isConnected)
	visited := make([]bool, cities)
	res := 0
	for i := 0; i < cities; i++ {
		if !visited[i] {
			dfs(isConnected, visited, cities, i)
			res++
		}
	}
	return res
}

func dfs(isConnected [][]int, visited []bool, cities int, i int) {
	for j := 0; j < cities; j++ {
		if isConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(isConnected, visited, cities, j)
		}
	}
}

//runtime 0ms
//Memory 8.52MB
