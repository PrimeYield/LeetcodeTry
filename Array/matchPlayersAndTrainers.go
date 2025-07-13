//2410. Maximum Matching of Players With Trainers

package array

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

//19ms Beats 83.72%
//10.71MB Beates 79.07%

// func matchPlayersAndTrainers(players []int, trainers []int) int {
// 	playersMap := make(map[int]struct{})
// 	trainersMap := make(map[int]struct{})
// 	res := 0
// 	for _, num := range players {
// 		playersMap[num] = struct{}{}
// 	}
// 	for _, num := range trainers {
// 		trainersMap[num] = struct{}{}
// 	}
// 	for i := 0; isEnd(playersMap, trainersMap); i++ {
// 		for num, _ := range playersMap {
// 			if _, ok := trainersMap[num+i]; ok {
// 				res++
// 				delete(playersMap, num)
// 				delete(trainersMap, (num + i))
// 			}
// 		}
// 	}
// 	return res
// }

// func isEnd(players, trainers map[int]struct{}) bool {
// 	trainerMax := 0
// 	playersMin := 0
// 	for num, _ := range trainers {
// 		trainerMax = max(trainerMax, num)
// 	}
// 	for num, _ := range players {
// 		playersMin = min(playersMin, num)
// 	}
// 	return trainerMax >= playersMin
// }
