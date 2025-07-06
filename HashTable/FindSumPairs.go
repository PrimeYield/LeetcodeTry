//1865. Finding Pairs With a Certain Sum

package hashtable

type FindSumPairs struct {
	mapNums1   map[int]int
	mapNums2   map[int]int
	arrayNums2 []int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	mapNums1, mapNums2 := make(map[int]int), make(map[int]int)
	for _, num := range nums1 {
		mapNums1[num]++
	}
	for _, num := range nums2 {
		mapNums2[num]++
	}
	return FindSumPairs{mapNums1: mapNums1, mapNums2: mapNums2, arrayNums2: nums2}
}

func (this *FindSumPairs) Add(index int, val int) {
	this.mapNums2[this.arrayNums2[index]]--
	this.arrayNums2[index] += val
	this.mapNums2[this.arrayNums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	res := 0
	for num, count := range this.mapNums1 {
		if _, exist := this.mapNums2[tot-num]; exist {
			res = res + (count * this.mapNums2[tot-num])
		}
	}
	return res
}
