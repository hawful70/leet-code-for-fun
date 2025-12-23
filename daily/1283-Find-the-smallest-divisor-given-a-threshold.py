from typing import List

class Solution:
    def smallestDivisor(self, nums: List[int], threshold: int) -> int:
        lo, hi = 1, max(nums)

        while lo < hi:
            mid = (lo + hi) // 2
            if self.valid_sum(nums, threshold, mid):
                hi = mid        # move LEFT → find first valid
            else:
                lo = mid + 1    # move RIGHT
        return lo

    def valid_sum(self, nums: List[int], threshold: int, divisor: int) -> bool:
        total = 0
        for num in nums:
            total += (num + divisor - 1) // divisor  # CEILING division
        return total <= threshold
