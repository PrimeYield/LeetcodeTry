// 347. Top K Frequent Elements
package heap

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	heap := make([]int, 1, len(countMap)+1)
	for key, _ := range countMap {
		heap = append(heap, key)
	}
	for i := len(heap) - 1; i/2 > 0; i-- {
		if countMap[heap[i]] > countMap[heap[i/2]] {
			heap[i], heap[i/2] = heap[i/2], heap[i]
		}
	}
	//目前只能確定heap[1]是最大值 heap[2]則未必
	res := make([]int, 0, k)
	for len(res) < k {
		res = append(res, heap[1])
		heap[1] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		for i := len(heap) - 1; i/2 > 0; i-- {
			if countMap[heap[i]] > countMap[heap[i/2]] {
				heap[i], heap[i/2] = heap[i/2], heap[i]
			}
		}
	}
	return res
}

//Runtime 35ms Beats 5.37%
//Memory 8.04MB Beats 45.55%
