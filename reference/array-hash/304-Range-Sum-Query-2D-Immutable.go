type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	sumMatrix := make([][]int, rows+1)
	for i := range sumMatrix {
		sumMatrix[i] = make([]int, cols+1)
	}

	for r := 0; r < rows; r++ {
		prefix := 0
		for c := 0; c < cols; c++ {
			prefix += matrix[r][c]
			above := sumMatrix[r][c+1]
			sumMatrix[r+1][c+1] = prefix + above
		}
	}
	return NumMatrix{sumMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1, col1, row2, col2 = row1+1, col1+1, row2+1, col2+1
	bottomRight := this.sumMatrix[row2][col2]
	above := this.sumMatrix[row1-1][col2]
	left := this.sumMatrix[row2][col1-1]
	topLeft := this.sumMatrix[row1-1][col1-1]
	return bottomRight - above - left + topLeft
}

/**
* Your NumMatrix object will be instantiated and called as such:
* obj := Constructor(matrix);
* param_1 := obj.SumRegion(row1,col1,row2,col2);
 */