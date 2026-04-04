package daily

func rotate(nums []int, k int) {
	if len(nums) > 1 {
		left, right := 0, len(nums)-1
		k = k % len(nums)
		reverseArray(left, right, nums)
		reverseArray(0, k-1, nums)
		reverseArray(k, right, nums)
	}
}

func reverseArray(left int, right int, nums []int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
