//118. Pascal's Triangle

package array

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	res = gen(numRows, res)
	return res
}

func gen(numRows int, res [][]int) [][]int {
	if numRows == 1 {
		res = [][]int{{1}}
		return res
	}
	if numRows == 2 { //numRows == n  => len = n
		res = [][]int{{1}, {1, 1}}
		return res
	}
	res = gen(numRows-1, res)
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
//4.51MB
