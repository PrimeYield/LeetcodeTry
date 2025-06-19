//724. Find Pivot Index

package prefixsum

func pivotIndex(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	arr := make([]int, len(nums)+1, len(nums)+1) //right
	arr[0] = total
	left := 0
	arrInd := 1
	for i := 0; i < len(nums); i++ {
		arr[arrInd] = arr[i] - nums[i]
		if i > 0 {
			left += nums[i-1]
		}
		if left == arr[arrInd] {
			return i
		}
		arrInd++
	}
	return -1
}

//runtime 0ms
//Memory 8.1MB
