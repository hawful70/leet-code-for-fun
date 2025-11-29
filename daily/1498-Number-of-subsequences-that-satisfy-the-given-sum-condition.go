package daily

import "sort"

func numSubseq(nums []int, target int) int {
	mod := 1000000007
	sort.Ints(nums)
	n := len(nums)

	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = (pow2[i-1] * 2) % mod
	}

	ans := 0
	for left := 0; left < n; left++ {
		limit := target - nums[left]
		if limit < nums[left] {
			break
		}

		lo, hi := left, n-1
		for lo <= hi {
			mid := (lo + hi) / 2
			if nums[mid] <= limit {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		right := hi
		if right >= left {
			ans = (ans + pow2[right-left]) % mod
		}
	}
	return ans
}
