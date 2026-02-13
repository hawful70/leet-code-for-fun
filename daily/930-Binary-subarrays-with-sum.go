package daily

func numSubarraysWithSum(nums []int, goal int) int {
	return atMost(nums, goal) - atMost(nums, goal-1)
}

func atMost(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}
	left, current_sum := 0, 0
	count := 0
	for right := 0; right < len(nums); right++ {
		current_sum += nums[right]
		for current_sum > goal {
			current_sum -= nums[left]
			left++
		}
		count += right - left + 1
	}
	return count
}
