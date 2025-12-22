class Solution:
    def maxValue(self, n: int, index: int, maxSum: int) -> int:
        # Check if a given peak x is feasible
        def feasible(x: int) -> bool:
            total = x
            # left side: index elements
            total += calcSide(x - 1, index)
            # right side: n - index - 1 elements
            total += calcSide(x - 1, n - index - 1)
            return total <= maxSum

        # Calculate the minimum sum of one side
        def calcSide(start: int, length: int) -> int:
            if length == 0:
                return 0

            # Case 1: we do NOT hit 1
            if start > length:
                end = start - length + 1
                return (start + end) * length // 2

            # Case 2: we hit 1
            # arithmetic part: start, start-1, ..., 1
            sum_desc = (start + 1) * start // 2
            # remaining positions filled with 1
            sum_flat = length - start
            return sum_desc + sum_flat

        # Binary search on the answer
        lo, hi = 1, maxSum
        ans = 1

        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if feasible(mid):
                ans = mid
                lo = mid + 1  # try bigger peak
            else:
                hi = mid - 1  # peak too large

        return ans
