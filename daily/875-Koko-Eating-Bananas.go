package daily

import "math"

func minEatingSpeed(piles []int, h int) int {
	max := piles[0]
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	left, right := 1, max
	ans := 0
	for left <= right {
		mid := (left + right) / 2
		if validSpeed(piles, h, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func validSpeed(piles []int, h int, speed int) bool {
	hours := 0
	for _, pile := range piles {
		hours += int(math.Ceil(float64(pile) / float64(speed)))
		if hours > h {
			return false
		}
	}
	return hours <= h
}
