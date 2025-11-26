class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        duplicate_character = set()
        left = 0
        ans = 0
        for right in range(len(s)):
            while s[right] in duplicate_character:
                duplicate_character.remove(s[left])
                left += 1
            duplicate_character.add(s[right])
            ans = max(ans, right - left + 1)
        return ans