package daily

func characterReplacement(s string, k int) int {
	n := len(s)
	left := 0
	res := 0
	seen := make(map[rune]int)
	maxFrequency := 0
	for right := 0; right < n; right++ {
		seen[rune(s[right])]++
		maxFrequency = max(maxFrequency, seen[rune(s[right])])
		for (right-left+1)-maxFrequency > k {
			seen[rune(s[left])]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
