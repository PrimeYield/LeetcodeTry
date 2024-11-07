//1929. Concatenation of Array

package array

func doubler(target []int) []int {
	target = append(target, target...)
	return target
}
func getConcatenation(nums []int) []int {
	return doubler(nums)
}
