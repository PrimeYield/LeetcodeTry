// 349. Intersection of Two Arrays
package array

/*
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	var newNums1, newNums2 []int
	countMapNums1 := make(map[int]int)
	countMapNums2 := make(map[int]int)
	for _, num := range nums1 {
		countMapNums1[num]++
	}
	for num, _ := range countMapNums1 {
		newNums1 = append(newNums1, num)
	}
	for _, num := range nums2 {
		countMapNums2[num]++
	}
	for num, _ := range countMapNums2 {
		newNums2 = append(newNums2, num)
	}
	for i := 0; i < len(newNums1); i++ {
		for j := 0; j < len(newNums2); j++ {
			if newNums1[i] == newNums2[j] {
				res = append(res, newNums1[i])
			}
		}
	}

	return res
}
*/

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				break
			}
		}
	}
	resExist := make(map[int]bool)
	for i := 0; i < len(res); i++ {
		resExist[res[i]] = true
	}
	res = []int{}
	for k, _ := range resExist {

		res = append(res, k)
	}
	return res
}
