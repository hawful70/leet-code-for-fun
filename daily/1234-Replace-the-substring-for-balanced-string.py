class Solution:
    def balancedString(self, s: str) -> int:
        n = len(s)
        limit = n // 4
        outside_frequency = defaultdict(int)
        for c in s:
            outside_frequency[c] += 1
        
        balanced = True
        for c in "QWER":
            if outside_frequency[c] > limit:
               balanced = False
        if balanced: return 0

        res = n
        left = 0
        for right in range(n):
            outside_frequency[s[right]] -= 1
            while left < n:
                is_fixable = True
                for c in "QWER":
                    if outside_frequency[c] > limit:
                        is_fixable = False
                        break
                if is_fixable:
                    res = min(res, right - left + 1)
                    outside_frequency[s[left]] += 1
                    left += 1
                else:
                    break
        return res


