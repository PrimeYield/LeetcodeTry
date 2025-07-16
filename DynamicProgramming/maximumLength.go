//3201. Find the Maximum Length of Valid Subsequence I

package dynamicprogramming

/*
func maximumLength(nums []int) int {
    for i:=0;i<len(nums);i++{
		nums[i] = nums[i] % 2
	}
	maxLenght := 0
	count := 0
	for i:=0;i<len(nums);i++{
		if nums[i] == 1 {
			count++
		}
	}
	maxLenght = max(maxLenght,count)
	count = 0
	for i:=0 ;i<len(nums);i++{
		if nums[i] == 0 {
			count++
		}
	}
	maxLenght = max(maxLenght,count)
	count = 0
	for i:=0;i<len(nums);i++{
		//0,1,0,1....
		if nums[i] == 0 && count % 2 == 0 {
			count++
		}else if nums[i] == 1 && count % 2 == 1 {
			count++
		}
	}
	maxLenght = max(maxLenght,count)
    count = 0
	for i:=0;i<len(nums);i++{
		if nums[i] == 1 && count % 2 == 0 {
			count++
		}else if nums[i] == 0 && count % 2 == 1{
			count++
		}
	}
	maxLenght = max(maxLenght,count)
	return maxLenght
}
//8ms 12.60MB
*/

func maximumLength(nums []int) int {
	count := [2]int{} //記錄奇、偶個數  = 純奇的長 及 純偶的長
	for i := 0; i < len(nums); i++ {
		count[nums[i]%2]++
	}
	//處理交替出現的 偶、奇 … 或 奇、偶…
	c := nums[0] % 2
	both := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == c {
			both++
			c = 1 - c
		}
	}
	return max(both, count[0], count[1])
}

//0ms 14.04MB
