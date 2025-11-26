package daily

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}

	maxDay := 0
	for _, day := range bloomDay {
		if day > maxDay {
			maxDay = day
		}
	}

	low, high := 1, maxDay
	ans := 0
	for low <= high {
		mid := (low + high) / 2
		if canMakeBouquets(bloomDay, m, k, mid) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func canMakeBouquets(bloomDay []int, m int, k int, targetDay int) bool {
	flowers := 0
	bouquets := 0
	for _, day := range bloomDay {
		if targetDay >= day {
			flowers++
			if flowers == k {
				bouquets++
				flowers = 0
			}
		} else {
			flowers = 0
		}
	}
	return bouquets >= m
}
