package daily

func numSubarraysWithSum(nums []int, goal int) int {
	return atMost(nums, goal) - atMost(nums, goal-1)
}
func atMost(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}
	left := 0
	n := len(nums)
	currentSum := 0
	counted := 0
	for right := 0; right < n; right++ {
		currentSum += nums[right]
		for currentSum > goal {
			currentSum -= nums[left]
			left++
		}
		counted += (right - left + 1)
	}
	return counted
}
