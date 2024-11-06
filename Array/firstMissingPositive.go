package array

func firstMissingPositive(nums []int) int {
	var res []int
	var ans int
	for i := 1; i < len(nums)+1; i++ {
		res = append(res, i)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		if nums[i] > len(nums) {
			continue
		}
		if nums[i] == res[nums[i]-1] {
			res[nums[i]-1] = 0
		}
	}
	for i := 0; i < len(res); i++ {
		if res[i] != 0 {
			ans = res[i]
			break
		} else {
			ans = len(nums) + 1
		}
	}
	return ans
}
