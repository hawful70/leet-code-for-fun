class Solution:
    def validPalindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1
        while left <= right:
            if s[left] != s[right]:
                return self.validPalindromeUtil(s, left + 1, right) or self.validPalindromeUtil(s, left, right - 1)
            else:
                left += 1
                right -= 1
        return True
        
    def validPalindromeUtil(self, s: str, l: int, r: int) -> bool:
        while l <= r:
            if s[l] != s[r]:
                return False
            else:
                l += 1
                r -= 1
        return True