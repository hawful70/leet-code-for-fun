package daily

func findBestValue(arr []int, target int) int {
	maxValue := 0
	for _, val := range arr {
		if val > maxValue {
			maxValue = val
		}
	}
	low, high := 1, maxValue
	for low < high {
		mid := (low + high) / 2
		if sumAtMid(arr, mid) <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	s1 := sumAtMid(arr, low)
	s2 := sumAtMid(arr, low-1)
	if abs(s1-target) < abs(s2-target) {
		return low
	}
	return low - 1
}

func sumAtMid(arr []int, value int) int {
	sum := 0
	for _, x := range arr {
		if x > value {
			sum += value
		} else {
			sum += x
		}
	}
	return sum
}
