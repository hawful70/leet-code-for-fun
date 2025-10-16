package daily

func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (high + low) / 2
		if nums[mid] < nums[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
