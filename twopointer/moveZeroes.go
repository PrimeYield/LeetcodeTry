//283. Move Zeroes

package twopointer

/*
func moveZeroes(nums []int) {
	for i:=0;i<len(nums);i++{
		if nums[i] == 0 {
			for j:=i+1;j<len(nums);j++{
				if nums[j] != 0 {
					nums[i],nums[j] = nums[j],nums[i]
                    break
				}
			}
		}
	}
}
//runtime 41ms
//Memory 8.6MB
*/

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums) && j < len(nums); j++ {
		if nums[i] != 0 {
			for ; i < j; i++ {
				if nums[i] == 0 {
					break
				}
			}
		}
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

//runtime 0ms
//Memory 9.6MB
