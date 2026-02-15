package daily

func equalSubstring(s string, t string, maxCost int) int {
	currentCost := 0
	left := 0
	n := len(s)
	maxLength := 0
	for right := 0; right < n; right++ {
		currentCost += abs(int(s[right]) - int(t[right]))
		for currentCost > maxCost {
			currentCost -= abs(int(s[left]) - int(t[left]))
			left++
		}
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
