// 347. Top K Frequent Elements
package array

// func topKFrequent(nums []int, k int) []int {
// 	var res []int
// 	countMap := make(map[int]int)
// 	for _, num := range nums {
// 		countMap[num]++
// 	}
// 	for i, _ := range countMap {
// 		res = append(res, i)
// 	}
// 	for i, j := 0, len(res)-1; i < j; j-- {
// 		if countMap[res[i]] < countMap[res[j]] {
// 			res[i], res[j] = res[j], res[i]
// 		}
// 		if j-i == 1 {
// 			i++
// 			j = len(res)
// 		}
// 	}
// 	return res[:k]
// }
// Runtime 89 ms
// Memory 6 MB
func topKFrequent(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}
	stay := make([][]int, len(nums))
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	if len(countMap) == k {
		var res []int
		for num, _ := range countMap {
			res = append(res, num)
		}
		return res
	}
	for num, count := range countMap {
		stay[count] = append(stay[count], num)
		//fmt.Println(stay[count], count)
	}
	var res []int
	for i := len(stay) - 1; i >= 0; i-- {
		for j := 0; j < len(stay[i]); j++ {
			//fmt.Println(i)
			if len(res) == k {
				break
			}
			res = append(res, stay[i][j])
		}
	}
	return res
}

//Runtime 30ms Beats 5.37%
//Memory 8.29MB Beats 18.15%
