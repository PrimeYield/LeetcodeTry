//3360. Stone Removal Game

package math

func canAliceWin(n int) bool {
	times := 0
	for i := 10; n >= i; i-- {
		n -= i
		times++
	}
	return times%2 == 1
}

//runtime 0ms
//Memory 3.85% Beats 61.90%
