package daily

func partitionLabels(s string) []int {
	lastIndex := map[rune]int{}
	for i := 0; i < len(s); i++ {
		lastIndex[rune(s[i])] = i
	}

	var res []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		end = max(end, lastIndex[rune(s[i])])
		if end == i {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}
