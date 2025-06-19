//643. Maximum Average Subarray I

package slidingwindow

/*
func findMaxAverage(nums []int, k int) float64 {
	sum := 0.0
	avg := 0.0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	avg = sum / float64(k)
	for i := 1; i <= len(nums)-k; i++ {
		tempSum := sum - float64(nums[i-1]) + float64(nums[i+k-1])
		sum = tempSum
		if avg > sum/float64(k) {
			continue
		}
		avg = sum / float64(k)
	}
	return avg
}
//tuntime 4ms
//Memory 10.2MB
*/

func findMaxAverage(nums []int, k int) float64 {
	sum := 0.0
	avg := 0.0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	avg = sum / float64(k)
	for i := 1; i <= len(nums)-k; i++ {
		tempSum := sum - float64(nums[i-1]) + float64(nums[i+k-1])
		sum = tempSum
		avg = max(avg, sum/float64(k))
	}
	return avg
}

//runtime 0ms
