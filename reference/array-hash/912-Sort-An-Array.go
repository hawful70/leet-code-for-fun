func sortArray(nums []int) []int {
	mergeSort(nums)
	return nums
}

func mergeSort(nums []int) {
	if len(nums) > 1 {
		// Split array into two halves
		mid := len(nums) / 2
		left := make([]int, mid)
		right := make([]int, len(nums)-mid)

		// Copy data to left and right slices
		copy(left, nums[:mid])
		copy(right, nums[mid:])

		// Recursively sort both halves
		mergeSort(left)
		mergeSort(right)

		// Merge the sorted halves
		i, j, k := 0, 0, 0
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				nums[k] = left[i]
				i++
			} else {
				nums[k] = right[j]
				j++
			}
			k++
		}

		// Copy remaining elements of left, if any
		for i < len(left) {
			nums[k] = left[i]
			i++
			k++
		}

		// Copy remaining elements of right, if any
		for j < len(right) {
			nums[k] = right[j]
			j++
			k++
		}
	}
}