package daily

func smallestDivisor(nums []int, threshold int) int {
	maxNum := 0
	for _, val := range nums {
		if val > maxNum {
			maxNum = val
		}
	}

	ans := 0
	lo, hi := 1, maxNum
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if validSum(nums, threshold, mid) {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func validSum(nums []int, threshold int, divisor int) bool {
	sum := 0
	for _, val := range nums {
		sum = sum + (val+divisor-1)/divisor
	}
	return sum <= threshold
}
