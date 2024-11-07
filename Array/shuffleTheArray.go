//1470. Shuffle the Array

package array

func shuffle(nums []int, n int) []int {
	res := make([]int, 2*n, 2*n)
	for i, j := 0, 0; i < n; i, j = i+1, j+2 {
		res[j], res[j+1] = nums[i], nums[i+n]
	}
	return res
}

/*
func shuffle(nums []int, n int) []int {
	numsX := nums[:n]
	numsY := nums[n:]
	res := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		res = append(res, numsX[i])
		res = append(res, numsY[i])
	}
	return res
}
*/
