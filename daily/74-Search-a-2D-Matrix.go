package daily

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		if matrix[i][cols-1] >= target {
			low, high := 0, cols-1
			for low <= high {
				mid := (low + high) / 2
				if matrix[i][mid] == target {
					return true
				} else if matrix[i][mid] > target {
					high = mid - 1
				} else {
					low = mid + 1
				}
			}
		}
	}
	return false
}
