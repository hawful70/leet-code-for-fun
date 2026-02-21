package daily

func balancedString(s string) int {
	n := len(s)
	limit := n / 4
	outsideFrequency := make(map[byte]int)
	for i := 0; i < n; i++ {
		outsideFrequency[s[i]]++
	}

	// Use a string to iterate through the target keys
	targetChars := "QWER"

	isBalanced := true
	for _, char := range targetChars {
		if outsideFrequency[byte(char)] > limit {
			isBalanced = false
			break
		}
	}
	if isBalanced {
		return 0
	}

	res := n
	left := 0
	for right := 0; right < n; right++ {
		outsideFrequency[s[right]]--

		// Manual "all" check
		for left <= right { // Optimized: left shouldn't pass right
			isFixable := true
			for _, char := range targetChars {
				if outsideFrequency[byte(char)] > limit {
					isFixable = false
					break
				}
			}

			if isFixable {
				if right-left+1 < res {
					res = right - left + 1
				}
				outsideFrequency[s[left]]++
				left++
			} else {
				break
			}
		}
	}
	return res
}
