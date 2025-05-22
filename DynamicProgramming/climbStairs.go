package dynamicprogramming

func climbStairs(n int) int {
	table := make(map[int]int)
	table[0] = 1
	table[1] = 1
	var dp func(n int) int
	dp = func(n int) int {
		if _, exists := table[n]; exists {
			return table[n]
		}
		table[n] = dp(n-1) + dp(n-2)
		return table[n]
	}
	return dp(n)
}

//runtime 0ms Memory 4.04MB
