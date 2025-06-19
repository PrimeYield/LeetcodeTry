//1732. Find the Highest Altitude

package prefixsum

func largestAltitude(gain []int) int {
	arr := make([]int, len(gain)+1, len(gain)+1)
	arrInd := 1
	gainInd := 0
	maxGain := arr[0]
	for gainInd < len(gain) {
		arr[arrInd] = arr[gainInd] + gain[gainInd]
		maxGain = max(maxGain, arr[arrInd])
		arrInd++
		gainInd++
	}
	return maxGain
}

//runtime 0ms
//Memory 4.2MB
