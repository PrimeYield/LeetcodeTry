//1498. Number of Subsequences That Satisfy the Given Sum Condition

package twopointer

import (
	"sort"
)

func numSubseq(nums []int, target int) int {
	res := 0
	f := make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		f[i] = f[i-1] * 2 % 1000000007
	}
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] <= target/2; i++ {
		k := sort.SearchInts(nums, (target-nums[i])+1) - 1
		res = (res + f[k-i]) % 1000000007
	}
	return res
}

//runtime 27ms Beats 85.29%
//Memory  12.37MB Beats 5.88%
