//42. Trapping Rain Water

package array

func trap(height []int) int {
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	sum := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				sum += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				sum += rightMax - height[right]
			}
			right--
		}
	}
	return sum
}

//runtime 0ms memory 8.35MB
