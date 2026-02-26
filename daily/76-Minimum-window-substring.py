class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if not s or not t:
            return ""
        
        ans = float('inf'), None, None
        left = 0

        formed = 0
        frequency_t = Counter(t)
        required = len(frequency_t)
        current_window = defaultdict(int)

        for right in range(len(s)):
            char_r = s[right]
            current_window[char_r] += 1

            if char_r in frequency_t and current_window[char_r] == frequency_t[char_r]:
                formed += 1
            
            while left <= right and formed == required:
                char_l = s[left]

                if right - left + 1 < ans[0]:
                    ans = right - left + 1, left, right
                current_window[char_l] -= 1
                if char_l in frequency_t and current_window[char_l] < frequency_t[char_l]:
                    formed -= 1
                left += 1
        return "" if ans[0] == float('inf') else s[ans[1] : ans[2] + 1]