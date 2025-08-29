//3021. Alice and Bob Playing Flower Game

package math

// func flowerGame(n int, m int) int64 {
// 	var res int64
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= m; j++ {
// 			if (i+j)%2 == 1 {
// 				res++
// 			}
// 		}
// 	}
// 	return res
// }
// //Time Limit Exceeded

func flowerGame(n int, m int) int64 {
	var res int64
	//n及m各有幾個奇數、偶數
	n1 := n % 2
	m1 := m % 2
	if n1 != 0 {
		n2 := n / 2
		n3 := (n / 2) + 1
		if m1 != 0 {
			m2 := m / 2
			m3 := (m / 2) + 1
			res = int64(n2*m3) + int64(n3*m2)
		} else if m1 == 0 {
			m2 := m / 2
			res = int64(n2*m2) + int64(n3*m2)
		}
	} else if n1 == 0 {
		n2 := n / 2
		if m1 != 0 {
			m2 := m / 2
			m3 := (m / 2) + 1
			res = int64(n2*m2) + int64(n2*m3)
		} else if m1 == 0 {
			m2 := m / 2
			res = int64(n2*m2) * 2
		}
	}
	return res
}
