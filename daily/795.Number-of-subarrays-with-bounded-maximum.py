from typing import List


class Solution:
    def numSubarrayBoundedMax(self, nums: List[int], left: int, right: int) -> int:
        return self._count(nums, right) - self._count(nums, left - 1)

    def _count(self, nums: List[int], bound: int) -> int:
        cur = 0
        total = 0
        for x in nums:
            if x <= bound:
                cur += 1
            else:
                cur = 0
            total += cur
        return total

