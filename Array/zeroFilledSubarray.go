//2348. Number of Zero-Filled Subarrays

package array

func zeroFilledSubarray(nums []int) int64 {
	var res int64
	for i := 0; i < len(nums); i++ {
		start := 0
		end := 0
		// length := start - end
		length := 0
		if nums[i] == 0 {
			start = i
			for i < len(nums) {
				if nums[i] != 0 {
					end = i - 1
					break
				}
				if i == len(nums)-1 {
					end = len(nums) - 1
					break
				}
				i++
			}
			length = end - start + 1
			res = res + int64(((1+length)*length)/2)
		}
		// fmt.Println(i)
		// fmt.Println(nums[i])
	}
	return res
}

//0ms
//9.46MB 100%
