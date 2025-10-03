package daily

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	score := 0
	sort.Ints(tokens)
	left, right := 0, len(tokens)-1

	if len(tokens) == 0 {
		return 0
	}

	for left < right {
		if score == 0 {
			if power >= tokens[left] {
				score += 1
				power = power - tokens[left]
			}
			left++
		} else {
			if power < tokens[left] {
				power = power + tokens[right]
				score -= 1
				right--
			} else {
				score += 1
				power = power - tokens[left]
				left++
			}
		}
	}

	if tokens[right] <= power {
		score++
	}

	return score
}
