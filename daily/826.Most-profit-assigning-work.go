package daily

import "sort"

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	jobs := make([][2]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [2]int{difficulty[i], profit[i]}
	}

	sort.Slice(jobs, func(a, b int) bool {
		return jobs[a][0] < jobs[b][0]
	})

	maxProfit := make([]int, n)
	maxProfit[0] = jobs[0][1]
	for i := 1; i < n; i++ {
		maxProfit[i] = max(maxProfit[i-1], jobs[i][1])
	}

	total := 0
	for _, capacity := range worker {
		index := findMostDifficultyofWorker(jobs, capacity)
		if index != -1 {
			total += maxProfit[index]
		}
	}
	return total
}

func findMostDifficultyofWorker(jobs [][2]int, capacity int) int {
	low, high := 0, len(jobs)-1
	ans := -1
	for low <= high {
		mid := (high + low) / 2
		if jobs[mid][0] <= capacity {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
