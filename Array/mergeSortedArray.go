// 88. Merge Sorted Array
package array

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	if len(nums1) > m {
// 		nums1 = nums1[:m]
// 	}
// 	if len(nums2) > n {
// 		nums2 = nums2[:n]
// 	}

// 	nums1 = append(nums1, nums2...)

// 	sort.Ints(nums1)
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	//nums1==0 代表nums1,nums2都是空切片
	//nums1即使沒有元素，也會包含nums2的長度(n)
	if len(nums2) == 0 {
		return
	}
	if len(nums1) == n {
		copy(nums1, nums2)
		return
	}
	//排除其1為0的情況
	copyNums1 := make([]int, m, m)
	copy(copyNums1, nums1[:m])
	index := 0
	for i, j := 0, 0; i < m || j < n; {
		if i == m {
			nums1[index] = nums2[j]
			j++
			index++
			continue
		}
		if j == n {
			nums1[index] = copyNums1[i]
			i++
			index++
			continue
		}
		if copyNums1[i] <= nums2[j] {
			nums1[index] = copyNums1[i]
			i++
			index++
			continue
		}
		if copyNums1[i] > nums2[j] {
			nums1[index] = nums2[j]
			j++
			index++
			continue
		}
	}
}

//Runtime 0ms Beats 100.00%
//Memory 4.25MB Beats 40.03%
