package array

/*
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	var median float64
	if len(nums1)%2 == 1 {
		sort.Ints(nums1)
		median = float64(nums1[(len(nums1)-1)/2])
	} else if len(nums1)%2 == 0 {
		sort.Ints(nums1)
		l := len(nums1) / 2
		sum := float64(nums1[l]) + float64(nums1[l-1])
		median = sum / 2
	}

	return median
}
*/
//失敗
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	length1 := len(nums1)
// 	lenght2 := len(nums2)
// 	i, j := 0, 0
// 	combin := make([]int, lenght2+length1)
// 	if length1 == 0 {
// 		combin = nums2
// 	}
// 	if lenght2 == 0 {
// 		combin = nums1
// 	}
// 	for length1 > 0 || lenght2 > 0 {
// 		if nums1[i] <= nums2[j] {
// 			combin[i+j] = nums1[i]
// 			length1--
// 		} else if nums1[i] > nums2[j] {
// 			combin[i+j] = nums2[j]
// 			lenght2--
// 		}
// 		if nums1[i] <= nums2[j] {
// 			i++
// 		} else {
// 			j++
// 		}
// 		if i == len(nums1) {
// 			i--
// 		} else if j == len(nums2) {
// 			j--
// 		}

// 	}
// 	mid := len(nums1) / 2
// 	var res float64
// 	if len(nums1)%2 == 0 {
// 		res = float64((nums1[mid] + nums1[mid-1]) / 2)
// 	} else {
// 		res = float64(nums1[mid])
// 	}
// 	return res
// }
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//組合同時排序
	i, j, k := 0, 0, 0
	combin := make([]int, len(nums1)+len(nums2))
	var res float64
	if len(nums1) == 0 {
		combin = nums2
	}
	if len(nums2) == 0 {
		combin = nums1
	}
	if len(nums1) != 0 && len(nums2) != 0 {
		for i+j < len(combin) {
			if nums1[i] <= nums2[j] {
				combin[k] = nums1[i]
				k++
				i++
				if i == len(nums1) {
					combin = combin[:k]
					combin = append(combin, nums2[j:]...)
					break
				}
				continue
			} else if nums1[i] > nums2[j] {
				combin[k] = nums2[j]
				k++
				j++
				if j == len(nums2) {
					combin = combin[:k]
					combin = append(combin, nums1[i:]...)
					break
				}
				continue
			}
		}
	}
	mid := len(combin) / 2
	if len(combin)%2 == 1 {
		res = float64(combin[mid])
	}
	if len(combin)%2 == 0 {
		res = float64(combin[mid]+combin[mid-1]) / 2
	}

	return res
}
