//739. Daily Temperatures

package monotonicstack

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	ans := make([]int, len(temperatures), len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)

	}
	return ans
}

//runtime 15ms Beats 46.40
//Memory 12.86MB
