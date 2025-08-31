from typing import List


class Solution:
    def diStringMatch(self, s: str) -> List[int]:
        n = len(s)
        low, high = 0, n
        res = [0] * (n + 1)
        for i, ch in enumerate(s):
            if ch == 'I':
                res[i] = low
                low += 1
            else:  # 'D'
                res[i] = high
                high -= 1
        res[n] = high if s[n - 1] == 'I' else low
        return res

