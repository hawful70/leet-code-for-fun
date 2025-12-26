package daily

func maxSideLength(mat [][]int, threshold int) int {
	rows, cols := len(mat), len(mat[0])

	prefix := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		prefix[i] = make([]int, cols+1)
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + mat[i-1][j-1]
		}
	}

	lo, hi := 0, min(rows, cols)
	for lo < hi {
		mid := (hi + lo + 1) / 2
		if validSquare(prefix, mid, threshold) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func validSquare(prefix [][]int, k int, threshold int) bool {
	if k == 0 {
		return true
	}

	rows, cols := len(prefix)-1, len(prefix[0])-1
	for i := 0; i+k <= rows; i++ {
		for j := 0; j+k <= cols; j++ {
			sum := prefix[i+k][j+k] - prefix[i][j+k] - prefix[i+k][j] + prefix[i][j]
			if sum <= threshold {
				return true
			}
		}
	}
	return false
}
