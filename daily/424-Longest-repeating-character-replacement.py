class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        n = len(s)
        left = 0
        res = -float('inf')
        seen = defaultdict(int)
        max_frequency = 0
        for right in range(n):
            seen[s[right]] += 1
            max_frequency = max(max_frequency, seen[s[right]])
            while (right - left + 1) - max_frequency > k:
                seen[s[left]] -= 1
                left += 1
            res = max(res, right - left + 1)


        return res if res != -float('inf') else 0