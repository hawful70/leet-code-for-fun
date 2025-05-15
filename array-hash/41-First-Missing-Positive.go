func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		index := nums[i] - 1
		for nums[i] > 0 && nums[i] < n && nums[index] != nums[i] {
			nums[index], nums[i] = nums[i], nums[index]
			index = nums[i] - 1
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}