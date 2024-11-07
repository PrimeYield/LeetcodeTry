// 136. Single Number
package array

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	fmt.Println(nums)
	var k int

	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			k = i
			break
		}
		k = i + 2
	}
	return nums[k]
}
