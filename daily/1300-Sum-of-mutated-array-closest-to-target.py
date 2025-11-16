from typing import List


class Solution:
    def findBestValue(self, arr: List[int], target: int) -> int:
        max_value = max(arr)
        low, high = 1, max_value

        while low < high:
            mid = (low + high) // 2
            if self._mutated_sum(arr, mid) <= target:
                low = mid + 1
            else:
                high = mid

        sum_low = self._mutated_sum(arr, low)
        sum_prev = self._mutated_sum(arr, low - 1)
        if abs(sum_low - target) < abs(sum_prev - target):
            return low
        return low - 1

    def _mutated_sum(self, arr: List[int], value: int) -> int:
        total = 0
        for num in arr:
            total += value if num > value else num
        return total
