//495. Teemo Attacking

package array

// func findPoisonedDuration(timeSeries []int, duration int) int {
// 	res := 0
// 	for i:=0;i<len(timeSeries);i++{
// 		if i<len(timeSeries)-1 && timeSeries[i+1] - timeSeries[i] >= duration {
// 			res += duration
// 		}else if i<len(timeSeries)-1 && timeSeries[i+1] - timeSeries[i] < duration {
// 			res+= timeSeries[i+1] - timeSeries[i]
// 		}
// 	}
// 	return res + duration
// }
// //runtime 2ms Beats 10%
// //Memory 8.99MB

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1]-timeSeries[i] >= duration {
			res += duration
		} else if timeSeries[i+1]-timeSeries[i] < duration {
			res += timeSeries[i+1] - timeSeries[i]
		}
	}
	return res + duration
}

//runtime 0ms
//Memory 8.94MB
