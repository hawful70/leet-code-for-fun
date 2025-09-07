from typing import List


class Solution:
    def sumZero(self, n: int) -> List[int]:
        res: List[int] = []
        if n % 2 == 1:
            res.append(0)
        for i in range(1, n // 2 + 1):
            res.extend([-i, i])
        return res

