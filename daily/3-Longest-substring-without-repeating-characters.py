class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        seen = set()
        left = 0
        res = 0
        n = len(s)
        for right in range(n):
            while s[right] in seen:
                seen.remove(s[left])
                left += 1
            res = max(res, right - left + 1)
            seen.add(s[right])
        return res