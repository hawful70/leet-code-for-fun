from typing import List


class Solution:
    def getNoZeroIntegers(self, n: int) -> List[int]:
        for a in range(1, n):
            b = n - a
            if not self._has_zero(a) and not self._has_zero(b):
                return [a, b]
        return []

    def _has_zero(self, x: int) -> bool:
        while x > 0:
            if x % 10 == 0:
                return True
            x //= 10
        return False

