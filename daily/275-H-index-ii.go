package daily

func hIndex(citations []int) int {
	n := len(citations)
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		h := n - mid
		if citations[mid] >= h {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return n - low
}
