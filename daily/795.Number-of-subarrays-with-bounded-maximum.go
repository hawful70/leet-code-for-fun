package daily

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	return count(nums, right) - count(nums, left-1)
}

func count(nums []int, bound int) int {
	cur := 0
	total := 0
	for _, num := range nums {
		if num <= bound {
			cur += 1
		} else {
			cur = 0
		}
		total += cur
	}
	return total
}
