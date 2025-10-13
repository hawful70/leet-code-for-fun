package daily

func searchMatrixII(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		low, high := 0, cols-1
		for low <= high {
			mid := (high + low) / 2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				high--
			} else {
				low++
			}
		}
	}
	return false
}
