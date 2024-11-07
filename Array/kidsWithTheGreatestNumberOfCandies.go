//1431. Kids With the Greatest Number of Candies

package array

import "sort"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	copyCandies := make([]int, len(candies))
	res := make([]bool, len(candies), len(candies))
	copy(copyCandies, candies)
	sort.Ints(copyCandies)
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= copyCandies[len(candies)-1] {
			res[i] = true
		}
	}
	return res
}
