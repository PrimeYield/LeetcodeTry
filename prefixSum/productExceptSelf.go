//238. Product of Array Except Self

package prefixsum

func productExceptSelf(nums []int) []int {
	length := len(nums)
	leftPre := make([]int, length)
	rightPre := make([]int, length)
	res := make([]int, length)
	leftPre[0] = 1
	rightPre[length-1] = 1
	for i := 1; i < length; i++ {
		leftPre[i] = leftPre[i-1] * nums[i-1]
	}
	for i := length - 2; i >= 0; i-- {
		rightPre[i] = rightPre[i+1] * nums[i+1]
	}
	for i := 0; i < length; i++ {
		res[i] = leftPre[i] * rightPre[i]
	}
	return res
}
