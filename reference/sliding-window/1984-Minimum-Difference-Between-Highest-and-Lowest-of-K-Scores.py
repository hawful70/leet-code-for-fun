class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        nums.sort()
        right = k - 1
        ans = float('inf')
        for left in range(len(nums)):
            if right < len(nums):
                ans = min(nums[right] - nums[left], ans)
                right += 1
        return ans