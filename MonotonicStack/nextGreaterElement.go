//496. Next Greater Element I

package monotonicstack

// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	length := len(nums1)
// 	res := make([]int, length)
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < len(nums2); j++ {
// 			if nums2[j] == nums1[i] {
// 				temp := nums2[j]
// 				for ; j < len(nums2); j++ {
// 					if temp < nums2[j] {
// 						res[i] = nums2[j]
// 						break
// 					} else {
// 						res[i] = -1
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// //runtime 1ms Beats 28.72%
// //Memory 4.62MB Beats 96.41%
// //沒有用到stack
// //O(m*n)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greaterMap := make(map[int]int)
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums2[i])
			continue
		}
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			greaterMap[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	res := make([]int, len(nums1), len(nums1))
	for i, num := range nums1 {
		if _, ok := greaterMap[num]; ok {
			res[i] = greaterMap[num]
		} else {
			res[i] = -1
		}
	}
	return res
}

//runtime 0ms
//Memory 5.22MB Beats 57.44%
