from typing import List


class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        if len(bloomDay) < m * k:
            return -1

        low, high = 1, max(bloomDay)
        ans = high

        while low <= high:
            mid = (low + high) // 2
            if self.can_make(bloomDay, m, k, mid):
                ans = mid
                high = mid - 1
            else:
                low = mid + 1

        return ans

    def can_make(self, bloomDay: List[int], m: int, k: int, day: int) -> bool:
        bouquets = flowers = 0
        for bloom in bloomDay:
            if bloom <= day:
                flowers += 1
                if flowers == k:
                    bouquets += 1
                    flowers = 0
            else:
                flowers = 0

        return bouquets >= m
