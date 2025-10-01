//1823. Find the Winner of the Circular Game

package queue

func findTheWinner(n int, k int) int {
	Q := make([]int, 0)
	for i := 1; i <= n; i++ {
		Q = append(Q, i)
	}
	temp := k
	for i := 0; len(Q) > 1; {
		if temp > 1 {
			temp--
			Q = append(Q, Q[i])
			Q = Q[1:]
		} else {
			temp = k
			Q = Q[1:]
		}
	}
	return Q[0]
}
