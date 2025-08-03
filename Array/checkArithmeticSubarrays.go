//1630. Arithmetic Subarrays

package array

// func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
// 	var res []bool
// 	m := len(l)
// 	dummy := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		dummy[i] = nums[i]
// 	}
// 	for i := 0; i < m; i++ {
// 		temp := make([]int, 0)
// 		for j := l[i]; j <= r[i]; j++ {
// 			temp = append(temp, dummy[j])
// 		}
// 		sort.Ints(temp)
// 		tempAri := temp[1] - temp[0]
// 		for j := 0; j < len(temp)-1; j++ {
// 			if temp[j+1]-temp[j] != tempAri {
// 				res = append(res, false)
// 				break
// 			}
// 			if j == len(temp)-2 {
// 				res = append(res, true)
// 			}
// 		}
// 	}
// 	return res
// }

// //45ms
// //9.46MB
func getMinMax(nums []int) (int, int) {
	low := nums[0]
	high := nums[0]
	for i := 1; i < len(nums); i++ {
		if low > nums[i] {
			low = nums[i]
		}
		if high < nums[i] {
			high = nums[i]
		}
	}
	return low, high
}

func isAtc(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	n := len(nums)
	low, high := getMinMax(nums)
	if (high-low)%(n-1) != 0 {
		return false
	}
	diff := (high - low) / (n - 1)
	vals := map[int]struct{}{}
	for i := 0; i < n; i++ {
		vals[nums[i]] = struct{}{}
	}
	for i := low + diff; i < high; i += diff {
		if _, ok := vals[i]; !ok {
			return false
		}
	}
	return true
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := []bool{}
	for i := 0; i < len(l); i++ {
		res = append(res, isAtc(nums[l[i]:r[i]+1]))
	}
	return res
}

//0ms
//8.71MB
