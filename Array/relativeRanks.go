//506. Relative Ranks

package array

import "strconv"

func findRelativeRanks(score []int) []string {
	var rank []int
	var res []string
	rank = append(rank, score...)
	for i, j := 0, len(rank)-1; i < j; j-- {
		if rank[i] < rank[j] {
			rank[i], rank[j] = rank[j], rank[i]
		}
		if j-i <= 1 {
			i++
			j = len(rank)
		}
	}
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(rank); j++ {
			if rank[j] == score[i] {
				if j+1 == 1 {
					res = append(res, "Gold Medal")
					break
				}
				if j+1 == 2 {
					res = append(res, "Silver Medal")
					break
				}
				if j+1 == 3 {
					res = append(res, "Bronze Medal")
					break
				}
				res = append(res, strconv.Itoa(j+1))
			}
		}
	}
	return res
}
