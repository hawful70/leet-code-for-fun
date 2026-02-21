package daily

func balancedString(s string) int {
	n := len(s)
	limit := n / 4
	outsideFrequency := make(map[byte]int)
	for i := 0; i < n; i++ {
		outsideFrequency[s[i]]++
	}

	balanced := true
	for i := 0; i < 4; i++ {
		if outsideFrequency[byte('Q'+i)] > limit {
			balanced = false
			break
		}
	}
	if balanced {
		return 0
	}

	res := n
	left := 0
	for right := 0; right < n; right++ {
		outsideFrequency[s[right]]--
		for left < n {
			isFixable := true
			for i := 0; i < 4; i++ {
				if outsideFrequency[byte('Q'+i)] > limit {
					isFixable = false
					break
				}
			}
			if isFixable {
				res = min(res, right-left+1)
				outsideFrequency[s[left]]++
				left++
			} else {
				break
			}
		}
	}
	return res
}
