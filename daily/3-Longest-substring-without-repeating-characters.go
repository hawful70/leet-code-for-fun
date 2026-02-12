package daily

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	left := 0
	res := 0
	seen := make(map[rune]int)
	for right := 0; right < n; right++ {
		seen[rune(s[right])]++
		for seen[rune(s[right])] > 1 {
			seen[rune(s[left])]--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}
