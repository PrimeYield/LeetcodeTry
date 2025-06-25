//649. Dota2 Senate

package queue

import "container/list"

// func predictPartyVictory(senate string) string {
// 	senateArray := []byte(senate)
// 	senatePower := make([]bool, len(senateArray), len(senateArray))
// 	for i := 0; i < len(senateArray); {
// 		if !senatePower[i] {
// 			for j := i; j < len(senateArray); j++ {
// 				if !senatePower[j] && senateArray[j] != senateArray[i] {
// 					senatePower[j] = !senatePower[j]
// 					senatePower = append(senatePower, senatePower[i])
// 					senatePower = senatePower[1:]
// 					senateArray = append(senateArray, senateArray[i])
// 					senateArray = senateArray[1:]
// 					break
// 				} else if j == len(senateArray)-1 && senateArray[j] == senateArray[i] {
// 					i = len(senateArray)
// 				}
// 			}
// 		} else {
// 			senatePower = senatePower[1:]
// 			senateArray = senateArray[1:]
// 		}
// 	}
// 	if senateArray[0] == 'R' {
// 		return "Radiant"
// 	}
// 	return "Dire"
// }
// //runtime 39ms
// //Memory 5.24MB

func predictPartyVictory(senate string) string {
	n := len(senate)
	RadQ := list.New()
	DirQ := list.New()

	for i, senator := range senate {
		if senator == 'R' {
			RadQ.PushBack(i)
		} else {
			DirQ.PushBack(i)
		}
	}

	for RadQ.Len() > 0 && DirQ.Len() > 0 {
		rIdx := RadQ.Remove(RadQ.Front()).(int)
		dIdx := DirQ.Remove(DirQ.Front()).(int)
		if rIdx > dIdx {
			DirQ.PushBack(dIdx + n)
		} else {
			RadQ.PushBack(rIdx + n)
		}
	}
	if RadQ.Len() > 0 {
		return "Radiant"
	}
	return "Dire"
}

//runtime 12ms Beats 16.42%
//Memory 8.27MB
