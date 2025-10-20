from typing import List


class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        n = len(matrix)
        low, high = matrix[0][0], matrix[-1][-1]

        def count_less_equal(target: int) -> int:
            row = n - 1
            col = 0
            count = 0

            while row >= 0 and col < n:
                if matrix[row][col] <= target:
                    count += row + 1
                    col += 1
                else:
                    row -= 1

            return count

        while low < high:
            mid = low + (high - low) // 2
            if count_less_equal(mid) < k:
                low = mid + 1
            else:
                high = mid

        return low
