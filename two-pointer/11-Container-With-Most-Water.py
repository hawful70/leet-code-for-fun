class Solution:
    def maxArea(self, height: List[int]) -> int:
        ans = 0
        left, right = 0, len(height) - 1
        while left < right:
            minHeight = min(height[left], height[right])
            current_container = minHeight * (right - left)
            ans = max(ans, current_container)
            left += 1
            right -= 1
        return ans