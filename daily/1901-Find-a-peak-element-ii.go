package daily

func findPeakGrid(mat [][]int) []int {
	rows, cols := len(mat), len(mat[0])
	lo, hi := 0, rows-1

	for lo <= hi {
		mid := (lo + hi) / 2

		maxCol := 0
		for c := 1; c < cols; c++ {
			if mat[mid][c] > mat[mid][maxCol] {
				maxCol = c
			}
		}

		currentElement := mat[mid][maxCol]
		up, down := 0, 0
		if mid > 0 {
			up = mat[mid-1][maxCol]
		}
		if mid < rows-1 {
			down = mat[mid+1][maxCol]
		}

		if mid > 0 && up > currentElement {
			hi = mid - 1
		} else if mid < rows-1 && down > currentElement {
			lo = mid + 1
		} else {
			return []int{mid, maxCol}
		}
	}
	return []int{-1, -1}
}
