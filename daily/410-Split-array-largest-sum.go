package daily

func splitArray(nums []int, k int) int {
	maxValue := 0
	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
	}

	sumArr := 0
	for _, num := range nums {
		sumArr += num
	}

	ans := 0
	lo, hi := maxValue, sumArr
	for lo <= hi {
		mid := (lo + hi) / 2
		if valueFit(nums, mid, k) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func valueFit(nums []int, limit int, k int) bool {
	count := 1
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > limit {
			count++
			sum = num
		}
	}
	return count <= k
}
