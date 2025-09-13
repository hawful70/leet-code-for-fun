package daily

func expressiveWords(s string, words []string) int {
	res := 0

	for _, word := range words {
		if isValidStretchy(s, word) {
			res++
		}
	}

	return res
}

func isValidStretchy(s string, w string) bool {
	if len(s) < len(w) {
		return false
	}

	i, j := 0, 0
	for i < len(s) && j < len(w) {
		countS, countW := 1, 1
		if s[i] == w[j] {
			for i+1 < len(s) && s[i+1] == s[i] {
				countS++
				i++
			}
			for j+1 < len(w) && w[j+1] == w[j] {
				countW++
				j++
			}

			if countS < 3 && countS != countW {
				return false
			}
			if countS >= 3 && countW > countS {
				return false
			}

			i++
			j++
		} else {
			return false
		}
	}
	return i == len(s) && j == len(w)
}
