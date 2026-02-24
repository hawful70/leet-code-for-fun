package daily

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	n := len(nums)
	maxDeque := []int{}

	for i := 0; i < n; i++ {
		if len(maxDeque) > 0 && i-maxDeque[0] >= k {
			maxDeque = maxDeque[1:]
		}
		for len(maxDeque) > 0 && nums[maxDeque[len(maxDeque)-1]] < nums[i] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, i)
		if i >= k-1 {
			res = append(res, nums[maxDeque[0]])
		}
	}
	return res
}
