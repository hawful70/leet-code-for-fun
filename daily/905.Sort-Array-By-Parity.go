package daily

func sortArrayByParity(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 0 {
			left++
		} else if nums[right]%2 != 0 {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return nums
}
