//2327. Number of People Aware of a Secret

package dynamicprogramming

// func PeopleAwareOfSecret(n int, delay int, forget int) int {
// 	const mod = 1000000007
// 	know := make([]int, n+1)
// 	share := make([]int, n+1)
// 	know[1] = 1
// 	for i := 2; i <= n; i++ {
// 		know[i] = (know[i-1] + know[max(i-delay, 0)]) % mod
// 		if i > forget {
// 			know[i] = (know[i-1] + know[max(i-delay, 0)] - know[max(i-forget, 0)]) % mod
// 		}
// 		if i > delay {
// 			share[i] = (know[i-delay]) % mod
// 			if i > forget {
// 				share[i] -= (know[i-forget]) % mod
// 			}
// 		}
// 	}
// 	return ((know[n] - know[n-forget]) + mod) % mod
// }

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const mod = 1000000007
	know := make([]int, n+1)
	share := make([]int, n+1)
	know[1] = 1
	for i := 2; i <= n; i++ {
		know[i] = (know[i-1] + know[max(i-delay, 0)]) % mod
		if i > forget {
			know[i] = (know[i-1] + know[max(i-delay, 0)] - know[max(i-forget, 0)]) % mod
		}
		if i > delay {
			share[i] = (know[i-delay]) % mod
			if i > forget {
				share[i] -= (know[i-forget]) % mod
			}
		}
	}
	return ((know[n]-know[max(n-forget, 0)])%mod + mod) % mod
}
