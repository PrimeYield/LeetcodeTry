//3487. Maximum Unique Subarray Sum After Deletion

package hashtable

func maxSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ele := make(map[int]struct{})
	ele_ := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := ele[nums[i]]; !ok && nums[i] > 0 {
			ele[nums[i]] = struct{}{}
		} else if _, ok = ele_[nums[i]]; !ok && nums[i] <= 0 {
			ele_[nums[i]] = struct{}{}
		}
	}
	sum := 0
	if len(ele) == 0 {
		max := -101
		for num, _ := range ele_ {
			if max < num {
				max = num
			}
		}
		return max
	}
	for num, _ := range ele {
		sum += num
	}
	return sum
}

//runtime 0ms
//Memory 4.36MB beats 52.78%
