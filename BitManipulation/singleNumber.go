//136. Single Number

package bitmanipulation

func singleNumber(nums []int) int {
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = cur ^ nums[i]
	}
	return cur
}

//runtime 0ms
//Memory  8.03MB
