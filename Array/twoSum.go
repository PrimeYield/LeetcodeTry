package array

func twoSum(nums []int, target int) []int {
	var res []int
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[target-nums[i]]; ok && numMap[target-nums[i]] != i {
			res = append(res, i, numMap[target-nums[i]])
			break
		}
	}
	return res
}

/*
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i,j}
			}
		}

	}

	return []int{}
}
*/
/*
func twoSum(nums []int, target int) []int {
	//var anynumber []int
	var answer []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				answer = []int{i,j}
			}
		}
	}
	fmt.Println(answer)
	return answer
}
*/
/*
func twoSum(nums []int, target int) []int {
	var anynumber []int
	var answer []int
	i := 0
	j := i + 1
	for i = 0; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				anynumber = append(anynumber, nums[i], nums[j])
				//fmt.Printf("%d %d\n", i, j)
				answer = append(answer, i, j)
			}
		}

	}
	fmt.Println(answer)
	return answer
}
*/
