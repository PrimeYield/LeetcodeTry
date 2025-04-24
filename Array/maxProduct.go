//152. Maximum Product Subarray

package array

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := (1 << 20) * -1
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			if nums[i] > max {
				return nums[i]
			}
		}
		multi := nums[i]
		for j := i + 1; j < len(nums); j++ {
			multi *= nums[j]
			if max < multi {
				max = multi
			}
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

//Runtime 276ms Beats 5.16%
//Memory 5.40MB Beats 57.17%
