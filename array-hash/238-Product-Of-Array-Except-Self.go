func productExceptSelf(nums []int) []int {
	product := make([]int, len(nums))
	prefixLeft, prefixRight := 1, 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixLeft *= 1
		} else {
			prefixLeft *= nums[i-1]
		}
		product[i] = prefixLeft
	}

	for i := len(nums) - 1; i > -1; i-- {
		if i == len(nums)-1 {
			prefixRight *= 1
		} else {
			prefixRight *= nums[i+1]
		}
		product[i] *= prefixRight
	}

	return product
}