//1672. Richest Customer Wealth

package array

func maximumWealth(accounts [][]int) int {
	var richest int
	for i := 0; i < len(accounts); i++ {
		var sum int
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if richest < sum {
			richest = sum
		}
	}
	return richest
}
