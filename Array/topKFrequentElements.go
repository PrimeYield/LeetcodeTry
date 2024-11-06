package array

func topKFrequent(nums []int, k int) []int {
	var res []int
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	for i, _ := range countMap {
		res = append(res, i)
	}
	for i, j := 0, len(res)-1; i < j; j-- {
		if countMap[res[i]] < countMap[res[j]] {
			res[i], res[j] = res[j], res[i]
		}
		if j-i == 1 {
			i++
			j = len(res)
		}
	}
	return res[:k]
}
