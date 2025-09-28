class Solution:
    def mySqrt(self, x: int) -> int:
        left, right = 1, x
        ans = 0
        while left <= right:
            mid = (left + right) // 2
            square = mid * mid
            if square == x:
                return mid
            elif square < x:
                ans = mid
                left = mid + 1
            else:
                right = mid - 1
        return ans