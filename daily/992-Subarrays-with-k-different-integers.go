package daily

func subarraysWithKDistinct(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
}

func atMost(nums []int, k int) int {
	count := make(map[int]int)
	left := 0
	res := 0
	n := len(nums)
	for right := 0; right < n; right++ {
		count[nums[right]]++
		for len(count) > k {
			count[nums[left]]--
			if count[nums[left]] == 0 {
				delete(count, nums[left])
			}
			left++
		}
		res += right - left + 1
	}
	return res
}
