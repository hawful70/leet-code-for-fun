from typing import List


class Solution:
    def countBinarySubstrings(self, s: str) -> int:
        n = len(s)
        left = 0
        right = 0
        ans = 0
        while left < n:
            while right < n and s[left] == s[right]:
                right += 1
            left_group = right - left

            start = right
            while right < n and s[start] == s[right]:
                right += 1
            right_group = right - start

            ans += min(left_group, right_group)
            left = start
        return ans

