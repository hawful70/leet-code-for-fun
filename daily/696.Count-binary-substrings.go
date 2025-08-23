package daily

func countBinarySubstrings(s string) int {
	left := 0
	right := 0
	n := len(s)
	result := 0

	for left < n {
		for right < n && s[left] == s[right] {
			right++
		}
		leftGroup := right - left

		start := right
		for right < n && s[start] == s[right] {
			right++
		}
		rightGroup := right - start

		result += min(leftGroup, rightGroup)
		left = start
	}

	return result
}
