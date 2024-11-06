package array

import "sort"

/*
func maximumGap(nums []int) int {
	var gap int
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > gap {
			gap = nums[i+1] - nums[i]
		}
	}
	return gap
}
未使用bucket sort
*/

func maximumGap(nums []int) int {
	var bucket [][]int
	var maxGap int
	bucketNum := 10
	for i := 0; i < bucketNum; i++ {
		bucket = append(bucket, []int{})
	}
	for i := 0; i < len(nums); i++ {
		idx := 0
		for j := 1; j < 1000000000; j *= 10 {
			if nums[i]/j >= 10 {
				idx++
			}
			if nums[i]/j < 10 {
				break
			}
		}
		bucket[idx] = append(bucket[idx], nums[i])
	}
	k := 0
	for i := 0; i < len(bucket); i++ {
		sort.Ints(bucket[i])
		for _, num := range bucket[i] {
			nums[k] = num
			k++
		}
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i]-nums[i-1] > maxGap {
			maxGap = nums[i] - nums[i-1]
		}
	}
	return maxGap
}
