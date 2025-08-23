package daily

import "math"

func judgeSquareSum(c int) bool {
	left := 0
	right := int(math.Ceil(math.Sqrt(float64(c))))

	for left <= right {
		result := left*left + right*right
		if result == c {
			return true
		} else if result < c {
			left++
		} else {
			right--
		}
	}
	return false
}
