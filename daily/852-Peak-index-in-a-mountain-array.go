package daily

func peakIndexInMountainArray(arr []int) int {
	low, high := 0, len(arr)-1
	for low < high {
		mid := (high + low) / 2
		if arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
