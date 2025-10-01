//1733. Minimum Number of People to Teach

package hashtable

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	cant := make(map[int]bool)
	for _, friendship := range friendships {
		canTalk := false
		lan := make(map[int]bool)
		//friendships[i][0] 掌握的語言
		for _, language := range languages[friendship[0]-1] {
			lan[language] = true
		}
		for _, language := range languages[friendship[1]-1] {
			if lan[language] {
				canTalk = true
				break
			}
		}
		if !canTalk {
			cant[friendship[0]-1] = true
			cant[friendship[1]-1] = true
		}
	}

	maxCnt := 0
	cnt := make([]int, n+1)
	for person := range cant {
		for _, lan := range languages[person] {
			cnt[lan]++
			maxCnt = max(maxCnt, cnt[lan])
		}
	}

	return len(cant) - maxCnt
}
