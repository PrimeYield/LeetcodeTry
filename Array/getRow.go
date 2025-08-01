//119. Pascal's Triangle II

package array

func getRow(rowIndex int) []int {
	temp := make([][]int, 0)
	// res := make([]int, 0)
	temp = gen2(rowIndex, temp)
	return temp[len(temp)-1]
}

func gen2(numRows int, res [][]int) [][]int {
	if numRows == 0 {
		res = [][]int{{1}}
		return res
	}
	if numRows == 1 { //numRows == n  => len = n
		res = [][]int{{1}, {1, 1}}
		return res
	}
	res = gen2(numRows-1, res)
	tempArr := make([]int, 0)
	tempArr = append(tempArr, 1)
	row := len(res)
	for i := 0; i < len(res[row-1])-1; i++ {
		temp := res[row-1][i] + res[row-1][i+1]
		tempArr = append(tempArr, temp)
	}
	tempArr = append(tempArr, 1)
	res = append(res, tempArr)
	return res
}

//0ms
//4.41MB
