package daily

func longestOnes(nums []int, k int) int {
	frequency_zero := make(map[int]int)
	count := 0
	left := 0
	n := len(nums)
	for right := 0; right < n; right++ {
		if nums[right] != 1 {
			frequency_zero[nums[right]]++
		}
		for frequency_zero[nums[right]] > k {
			if nums[left] != 1 {
				frequency_zero[nums[left]]--
			}
			left++
		}
		count = max(count, right-left+1)
	}
	return count
}
