package daily

func firstBadVersion(n int) int {
	low, high := 1, n
	ans := 0
	for low <= high {
		mid := (high + low) / 2
		if isBadVersion(mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
