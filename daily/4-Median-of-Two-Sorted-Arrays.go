package daily

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	leftSize := (m + n + 1) / 2

	lo, hi := 0, m
	for lo <= hi {
		mid := (lo + hi) / 2
		mid2 := leftSize - mid

		var left1 int
		var right1 int
		var left2 int
		var right2 int

		if mid == 0 {
			left1 = -1 << 60
		} else {
			left1 = nums1[mid-1]
		}

		if mid2 == 0 {
			left2 = -1 << 60
		} else {
			left2 = nums2[mid2-1]
		}

		if mid == m {
			right1 = 1 << 60
		} else {
			right1 = nums1[mid]
		}

		if mid2 == n {
			right2 = 1 << 60
		} else {
			right2 = nums2[mid2]
		}

		if left1 <= right2 && left2 <= right1 {
			if (m+n)%2 == 1 {
				return float64(max(left1, left2))
			}

			maxLeft := max(left1, left2)
			minRight := min(right1, right2)
			return float64(maxLeft+minRight) / 2.0
		}

		if left1 > right2 {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return 0.0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
