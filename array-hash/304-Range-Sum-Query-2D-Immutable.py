class NumMatrix:
    def __init__(self, matrix: list[list[int]]):
        rows, cols = len(matrix), len(matrix[0])
        
        # Create sumMat with an extra row and column filled with 0
        self.sumMat = [[0] * (cols + 1) for _ in range(rows + 1)]

        # Compute prefix sum
        for r in range(rows):
            prefix = 0
            for c in range(cols):
                prefix += matrix[r][c]  # Row-wise prefix sum
                above = self.sumMat[r][c + 1]
                self.sumMat[r + 1][c + 1] = prefix + above

    def sumRegion(self, row1: int, col1: int, row2: int, col2: int) -> int:
        # Convert matrix indices to prefix sum indices
        row1, col1, row2, col2 = row1 + 1, col1 + 1, row2 + 1, col2 + 1

        bottomRight = self.sumMat[row2][col2]
        above = self.sumMat[row1 - 1][col2]
        left = self.sumMat[row2][col1 - 1]
        topLeft = self.sumMat[row1 - 1][col1 - 1]

        return bottomRight - above - left + topLeft
