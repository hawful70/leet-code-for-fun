func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	n := len(nums)
	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}

	return ans
}