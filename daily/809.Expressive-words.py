from typing import List


class Solution:
    def expressiveWords(self, s: str, words: List[str]) -> int:
        ans = 0
        for w in words:
            if self._is_stretchy(s, w):
                ans += 1
        return ans

    def _is_stretchy(self, s: str, w: str) -> bool:
        if len(s) < len(w):
            return False

        i = j = 0
        n, m = len(s), len(w)
        while i < n and j < m:
            if s[i] != w[j]:
                return False

            count_s = 1
            while i + 1 < n and s[i + 1] == s[i]:
                count_s += 1
                i += 1

            count_w = 1
            while j + 1 < m and w[j + 1] == w[j]:
                count_w += 1
                j += 1

            # If the group in s is less than 3, counts must match exactly.
            if count_s < 3 and count_s != count_w:
                return False
            # If the group in s is 3 or more, the word group cannot exceed it.
            if count_s >= 3 and count_w > count_s:
                return False

            i += 1
            j += 1

        return i == n and j == m

