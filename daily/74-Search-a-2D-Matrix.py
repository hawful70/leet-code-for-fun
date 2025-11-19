class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        rows, cols = len(matrix), len(matrix[0])
        for r in range(rows):
            if matrix[r][cols - 1] >= target:
                left, right = 0, cols - 1
                while left <= right:
                    mid = (left + right) // 2
                    if matrix[r][mid] == target:
                        return True
                    elif matrix[r][mid] > target:
                        right = mid - 1
                    else:
                        left = mid + 1
        return False
