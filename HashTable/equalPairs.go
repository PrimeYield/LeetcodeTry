//2352. Equal Row and Column Pairs

package hashtable

import "fmt"

/*
func equalPairs(grid [][]int) int {
	length := len(grid)
	row := make(map[string]int)
	for i := 0; i < length; i++ {
		temp := ""
		for j := 0; j < length; j++ {
			temp += fmt.Sprint(grid[i][j]) +","
		}
		row[temp]++
	}
	res:=0
	for i :=0;i<length;i++{
		temp := ""
		for j:=0;j<length;j++{
			temp += fmt.Sprint(grid[j][i]) +","
		}
		if _,exists := row[temp]; exists {
			res += row[temp]
		}
	}
	return res
}
//runtime 352ms
//Memory 28.8MB
*/

func equalPairs(grid [][]int) int {
	row := make(map[string]int)
	for _, rowArray := range grid {
		row[fmt.Sprint(rowArray)]++
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		temp := make([]int, 0)
		for j := 0; j < len(grid); j++ {
			temp = append(temp, grid[j][i])
		}
		if val, ok := row[fmt.Sprint(temp)]; ok {
			res += val
		}
	}
	return res
}

//runtime 52ms
//Memory 9.2MB
