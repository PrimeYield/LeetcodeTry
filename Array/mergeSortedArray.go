//88. Merge Sorted Array
package array

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) > m {
		nums1 = nums1[:m]
	}
	if len(nums2) > n {
		nums2 = nums2[:n]
	}

	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)
}
