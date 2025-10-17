package daily

import "sort"

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	start_interval := make([][2]int, n)
	for i := 0; i < n; i++ {
		start_interval[i] = [2]int{intervals[i][0], i}
	}

	sort.Slice(start_interval, func(a, b int) bool {
		return start_interval[a][0] < start_interval[b][0]
	})

	ans := make([]int, n)
	for i, interval := range intervals {
		index := findingRightInterval(start_interval, interval[1])
		if index != -1 {
			ans[i] = start_interval[index][1]
		} else {
			ans[i] = index
		}
	}

	return ans
}

func findingRightInterval(nums [][2]int, target int) int {
	low, high := 0, len(nums)-1
	ans := -1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid][0] >= target {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
