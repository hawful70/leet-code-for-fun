from collections import Counter
from typing import List


class Solution:
    def threeSumMulti(self, arr: List[int], target: int) -> int:
        mod = 10**9 + 7
        counts = Counter(arr)
        keys = sorted(counts)

        ans = 0
        for i, x in enumerate(keys):
            for y in keys[i:]:
                z = target - x - y
                if z < y:
                    continue
                if z not in counts:
                    continue

                cx, cy, cz = counts[x], counts[y], counts[z]

                if x == y == z:
                    ans += self._comb3(cx)
                elif x == y:
                    ans += self._comb2(cx) * cz
                elif y == z:
                    ans += cx * self._comb2(cy)
                else:
                    ans += cx * cy * cz

        return ans % mod

    @staticmethod
    def _comb2(n: int) -> int:
        return n * (n - 1) // 2

    @staticmethod
    def _comb3(n: int) -> int:
        return n * (n - 1) * (n - 2) // 6

