package daily

func numberOfSubarrays(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
}

func atMost(nums []int, k int) int {
	left, count, res := 0, 0, 0
	for right := range len(nums) {
		if nums[right]%2 != 0 {
			count += 1
		}
		for count > k {
			if nums[left]%2 != 0 {
				count -= 1
			}
			left++
		}
		res += right - left + 1
	}
	return res

}
