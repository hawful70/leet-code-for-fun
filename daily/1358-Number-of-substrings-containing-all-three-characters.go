package daily

func numberOfSubstrings(s string) int {
	n := len(s)
	left := 0
	res := 0
	count := map[byte]int{'a': 0, 'b': 0, 'c': 0}

	for right := 0; right < n; right++ {
		count[s[right]]++
		for count['a'] > 0 && count['b'] > 0 && count['c'] > 0 {
			res += n - right
			count[s[left]]--
			left++
		}
	}
	return res
}
