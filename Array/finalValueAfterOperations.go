//2011. Final Value of Variable After Performing Operations

package array

/*
func finalValueAfterOperations(operations []string) int {
	rule := make(map[string]int)
	var sum int
	rule = map[string]int{
		"++X": 1,
		"X++": 1,
		"--X": -1,
		"X--": -1,
	}
	for i := 0; i < len(operations); i++ {
		sum = sum + rule[operations[i]]
	}
	return sum
}
*/

func finalValueAfterOperations(operations []string) int {
	var sum int
	for i := 0; i < len(operations); i++ {
		if operations[i] == "++X" || operations[i] == "X++" {
			sum++
		} else {
			sum--
		}
	}
	return sum
}
