// 134. Gas Station
package dynamicprogramming

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	curGas := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		totalGas = totalGas + gas[i] - cost[i]
		curGas = curGas + gas[i] - cost[i]
		if curGas < 0 {
			curGas = 0
			start = i + 1
		}
	}
	if totalGas < 0 {
		return -1
	}
	return start
}

//runtime 0ms
//Memory 11.44MB
