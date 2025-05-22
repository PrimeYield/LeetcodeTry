package dynamicprogramming

// func fib(n int) int {
//     if n == 0 {
// 		return 0
// 	} else if n > 0 && n < 2 {
// 		return 1
// 	}

// 	sum := 0
// 	sum = fib(n-1) + fib(n-2)
// 	return sum
// }
// //runtime 11ms Memory 3.80MB

func fib(n int) int {
	dp := make([]int, n, n)
	if n < 2 {
		return n
	}
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

//runtime 1ms Memory 3.86MB
