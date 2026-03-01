package daily

import "math"

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	res := math.MaxInt32
	// Prefix sum array: P[i] is sum of first i elements
	p := make([]int, n+1)
	for i := 0; i < n; i++ {
		p[i+1] = p[i] + nums[i]
	}

	dq := []int{} // Will store indices in increasing order of P[i]

	for i := 0; i < n+1; i++ {
		// 1. OPTIMIZE: If we found a valid subarray, try to shrink from the left
		for len(dq) > 0 && p[i]-p[dq[0]] >= k {
			res = min(res, i-dq[0])
			dq = dq[1:]
		}

		// 2. MAINTAIN MONOTONICITY: If current P[i] is smaller than the last one,
		// the last one is useless because P[i] is a "better" starting point (smaller P[i] makes P[j]-P[i] larger)
		for len(dq) > 0 && p[i] <= p[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, i)
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}
