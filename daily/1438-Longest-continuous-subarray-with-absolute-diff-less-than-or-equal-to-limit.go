package daily

func longestSubarray(nums []int, limit int) int {
	maxDeque := []int{}
	minDeque := []int{}

	n := len(nums)
	left := 0
	res := 0

	for right := 0; right < n; right++ {
		// Maintain maxDeque (decreasing)
		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] < nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, right)

		// Maintain minDeque (increasing)
		for len(minDeque) > 0 && nums[minDeque[len(minDeque)-1]] > nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, right)

		// Shrink window if condition violated
		for nums[maxDeque[0]]-nums[minDeque[0]] > limit {
			left++
			if maxDeque[0] < left {
				maxDeque = maxDeque[1:]
			}
			if minDeque[0] < left {
				minDeque = minDeque[1:]
			}
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
