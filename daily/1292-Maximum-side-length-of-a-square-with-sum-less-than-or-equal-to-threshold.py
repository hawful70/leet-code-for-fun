from typing import List

class Solution:
    def maxSideLength(self, mat: List[List[int]], threshold: int) -> int:
        rows, cols = len(mat), len(mat[0])

        # Build prefix sum (rows+1 x cols+1)
        prefix = [[0] * (cols + 1) for _ in range(rows + 1)]

        for i in range(1, rows + 1):
            for j in range(1, cols + 1):
                prefix[i][j] = (
                    prefix[i - 1][j]
                    + prefix[i][j - 1]
                    - prefix[i - 1][j - 1]
                    + mat[i - 1][j - 1]
                )

        # Upper-bound binary search on k
        lo, hi = 0, min(rows, cols)

        while lo < hi:
            mid = lo + (hi - lo + 1) // 2  # bias right
            if self.valid_square(prefix, mid, threshold):
                lo = mid
            else:
                hi = mid - 1

        return lo

    def valid_square(self, prefix: List[List[int]], k: int, threshold: int) -> bool:
        if k == 0:
            return True

        rows, cols = len(prefix) - 1, len(prefix[0]) - 1

        for i in range(rows - k + 1):
            for j in range(cols - k + 1):
                square_sum = (
                    prefix[i + k][j + k]
                    - prefix[i][j + k]
                    - prefix[i + k][j]
                    + prefix[i][j]
                )
                if square_sum <= threshold:
                    return True

        return False
