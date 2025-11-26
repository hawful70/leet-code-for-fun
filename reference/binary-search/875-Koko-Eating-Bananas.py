class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        left, right = 1, max(piles)
        ans = 0
        while left <= right:
            mid = (left + right) // 2
            if self.validSpeed(piles, h, mid):
                ans = mid
                right = mid - 1
            else:
                left = mid + 1
        return ans
    def validSpeed(self, piles: List[int], h: int, speed: int) -> bool:
        hours = 0
        for pile in piles:
            hours += math.ceil(float(pile) / speed)
            if hours > h:
                return False
        return hours <= h