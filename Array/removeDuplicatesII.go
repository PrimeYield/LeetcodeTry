// 80. Remove Duplicates from Sorted Array II
package array

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) int {
	countMap := make(map[int]int)
	var k int
	for i, num := range nums {
		countMap[num]++
		fmt.Println(countMap)
		if countMap[num] > 2 {
			nums[i] = 999999999
			k++
		}
	}
	sort.Ints(nums)
	nums = nums[:len(nums)-k]
	return len(nums)
}
