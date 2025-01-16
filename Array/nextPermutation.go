//31. Next Permutation

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i := len(nums) - 1
	j := len(nums) - 2
	k := len(nums) - 1
	for ; ; i, j = i-1, j-1 {
		if nums[i] > nums[j] {
			fmt.Println(j)
			break
		}
		if j == 0 && nums[i] < nums[j] {
			sort.Ints(nums)
			return
		}
		if j == 0 {
			break
		}
	}
	// if nums[k] <= nums[j] {
	// 	k--
	// }
	for {
		if nums[k] <= nums[j] {
			k--
		} else {
			break
		}
	}
	fmt.Println(k)
	nums[j], nums[k] = nums[k], nums[j]
	//sort[j+1:]
	sort.Ints(nums[j+1:])

}

//Runtime 0ms Beats 100.00%
//Memory 4.29 MB Beats 71.11%