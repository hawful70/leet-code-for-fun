import math


class Solution:
    def judgeSquareSum(self, c: int) -> bool:
        left = 0
        right = int(math.ceil(math.sqrt(c)))
        while left <= right:
            s = left * left + right * right
            if s == c:
                return True
            if s < c:
                left += 1
            else:
                right -= 1
        return False

