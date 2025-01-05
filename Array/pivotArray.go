//2161. Partition Array According to Given Pivot

package array

func pivotArray(nums []int, pivot int) []int {
	var res, smallers, biggers, pivots []int
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			smallers = append(smallers, nums[i])
		}
		if nums[i] > pivot {
			biggers = append(biggers, nums[i])
		}
		if nums[i] == pivot {
			pivots = append(pivots, nums[i])
		}
	}
	res = append(res, smallers...)
	res = append(res, pivots...)
	res = append(res, biggers...)
	return res
}

// Runtime 33ms Beats 10.26%
// Memory 11.24MB Beats 74.36%

// func pivotArray(nums []int, pivot int) []int {
// 	pivotIndex := findPivotIndex(nums, pivot)
// sortArray:
// 	for i := 0; i < len(nums); {
// 		if nums[i] > pivot && i < pivotIndex {
// 			nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
// 			pivotIndex = i
// 		}
// 		if nums[i] < pivot && i > pivotIndex {
// 			nums[i], nums[pivotIndex] = nums[pivotIndex], nums[i]
// 			pivotIndex = i
// 		}
// 		i++
// 	}
// 	if !checkSort(nums) {
// 		goto sortArray
// 	}
// 	return nums
// }

// func findPivotIndex(nums []int, pivot int) int {
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		if nums[i] == pivot {
// 			return i
// 		}
// 	}
// 	return len(nums) - 1
// }
// func checkSort(nums []int) bool {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > nums[i+1] {
// 			return false
// 		}
// 	}
// 	return true
// }
// Time Limit Exceeded

/*變成quicksort
func pivotArray(nums []int, pivot int) []int {
	if checkSort(nums) {
		return nums
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
	return pivotArray(nums, pivot)
}
func checkSort(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
*/

// func pivotArray(nums []int, pivot int) []int {
// 	if checkSort(nums, pivot) {
// 		return nums
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > pivot {
// 			if i < findPivot(nums, pivot) {
// 				for j := i; j < findPivot(nums, pivot); j++ {
// 					nums[j], nums[j+1] = nums[j+1], nums[j]
// 				}
// 			}
// 		}
// 		if nums[i] <= pivot {
// 			if i > findPivot(nums, pivot) {
// 				for j := i; j > findPivot(nums, pivot); j-- {
// 					nums[j], nums[j-1] = nums[j-1], nums[j]
// 				}
// 			}
// 		}
// 	}
// 	return pivotArray(nums, pivot)
// }

// func checkSort(nums []int, pivot int) bool {
// 	pivotPoint := findPivot(nums, pivot)
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > pivot && i < pivotPoint {
// 			return false
// 		}
// 		if nums[i] <= pivot && i > pivotPoint {
// 			return false
// 		}
// 	}
// 	return true
// }
// func findPivot(nums []int, pivot int) int {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == pivot {
// 			return i
// 		}
// 	}
// 	return 0
// }
