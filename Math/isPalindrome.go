package math

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	stringx := strconv.Itoa(x)
	runex := []rune(stringx)

	var Rerunex []rune

	for i := len(runex) - 1; i >= 0; i-- {
		Rerunex = append(Rerunex, runex[i])

	}
	ReStringx := string(Rerunex)
	if ReStringx == stringx {
		fmt.Println(x)
		return true
	} else {
		fmt.Printf("From left to right, it reads %s. From right to left, it becomes %s. Therefore it is not a palindrome.\n", stringx, ReStringx)
		return false

	}
	//fmt.Println(Rerunex) ok

}
