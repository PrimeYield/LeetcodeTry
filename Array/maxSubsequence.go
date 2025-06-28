//2099. Find Subsequence of Length K With the Largest Sum

package array

import "sort"

func maxSubsequence(nums []int, k int) []int {
	idx := make(map[int][]int)
	for i, num := range nums {
		idx[num] = append(idx[num], i)
	}
	dummy := make([]int, len(nums))
	_ = copy(dummy, nums)
	sort.Ints(nums)
	res := make([]int, 0)
	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		res = append(res, idx[nums[i]][0])
		if len(idx[nums[i]]) > 1 {
			idx[nums[i]] = idx[nums[i]][1:]
		}
	}
	sort.Ints(res)
	for i := 0; i < k; i++ {
		res[i] = dummy[res[i]]
	}
	return res
}

//runtime 0ms
//Memory 6MB
