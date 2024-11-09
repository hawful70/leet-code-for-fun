class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        repeating_character = {}
        ans = 0
        left = 0
        maxR = 0

        for right, c in enumerate(s):
            if c not in repeating_character:
                repeating_character[c] = 1
            else:
                repeating_character[c] += 1
            maxR = max(maxR, repeating_character[c])

            if right - left + 1 - maxR > k:
                repeating_character[s[left]] -= 1
                left += 1
                maxR -= 1
            
            ans = max(ans, right - left + 1)
        return ans