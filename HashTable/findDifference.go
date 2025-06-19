//2215. Find the Difference of Two Arrays

package hashtable

/*
func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]] = true
	}
	var res [][]int
	temp1 := make([]int, 0)
	temp2 := make([]int, 0)
	for _, num := range nums1 {
		if !map2[num] {
			temp1 = append(temp1, num)
			map2[num] = true
		}
	}
	res = append(res, temp1)
	for _, num := range nums2 {
		if !map1[num] {
			temp2 = append(temp2, num)
			map1[num] = true
		}
	}
	res = append(res, temp2)
	return res
}
//runtime 12ms
//Memory 9.7MB
*/

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]] = true
	}
	map2 := make(map[int]bool)
	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]] = true
	}
	temp1 := make([]int, 0)
	temp2 := make([]int, 0)
	for num, _ := range map1 {
		if !map2[num] {
			temp1 = append(temp1, num)
		} else {
			delete(map2, num)
		}
	}
	for num, _ := range map2 {
		temp2 = append(temp2, num)
	}
	var res [][]int
	res = append(res, temp1)
	res = append(res, temp2)
	return res
}

//runtime 4ms Beats: 88.47%
//Memory 9.60MB
