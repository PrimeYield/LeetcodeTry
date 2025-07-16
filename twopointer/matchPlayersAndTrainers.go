//2410. Maximum Matching of Players With Trainers

package twopointer

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	// hasPlayer := make([]bool, len(trainers), len(trainers))
	res := 0
	for i := len(trainers) - 1; len(players) > 0 && i >= 0; {
		if players[len(players)-1] <= trainers[i] {
			res++
			players = players[:len(players)-1]
			i--
			continue
		}
		players = players[:len(players)-1]
	}
	return res
}

//19ms Beats 91.69%
//10.71MB  Beats 82.53%
