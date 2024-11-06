package array

/*
func intersect(nums1 []int, nums2 []int) []int {
	countMapNums1 := make(map[int]int)
	countMapNums2 := make(map[int]int)
	var res []int
	for _, num := range nums1 {
		countMapNums1[num]++
	}
	for _, num := range nums2 {
		countMapNums2[num]++
	}
	for i, count := range countMapNums1 {
		if _, ok := countMapNums2[i]; ok {
			k := count
			for k > 1 {
				if k > countMapNums2[i] {
					k--
					continue
				}
				res = append(res, i)
				k--
			}
			res = append(res, i)
		}
	}
	return res
}
*/
func intersect(nums1 []int, nums2 []int) []int {

}
