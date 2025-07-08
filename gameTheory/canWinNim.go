//292. Nim Game

package gametheory

func canWinNim(n int) bool {
	//要保證取完能剩4的倍數
	if (n-1)%4 == 0 || (n-2)%4 == 0 || (n-3)%4 == 0 {
		return true
	}
	return false
}

//0ms
//3.93MB
