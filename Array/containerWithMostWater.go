// 11. Container With Most Water
package array

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	var area, maxArea int
	for left < right {
		l := right - left
		if height[left] > height[right] {
			area = height[right] * l
			right--
		} else {
			area = height[left] * l
			left++
		}
		if maxArea > area {
			maxArea = maxArea
		} else {
			maxArea = area
		}
	}
	return maxArea
}
