package daily

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	low, high := matrix[0][0], matrix[n-1][n-1]

	for low < high {
		mid := low + (high-low)/2
		if findingSmallerElementAtMid(matrix, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func findingSmallerElementAtMid(matrix [][]int, target int) int {
	n := len(matrix)
	row := n - 1
	col := 0
	count := 0

	for row >= 0 && col < n {
		if matrix[row][col] <= target {
			count += row + 1
			col++
		} else {
			row--
		}
	}

	return count
}
