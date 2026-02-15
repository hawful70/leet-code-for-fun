class Solution:
    def numSubarraysWithSum(self, nums: List[int], goal: int) -> int:
        return self.atMost(nums, goal) - self.atMost(nums, goal - 1)
    def atMost(self, nums: List[int], goal: int) -> int:
        if goal < 0:
            return 0
        left = 0
        n = len(nums)
        currentSum = 0
        counted = 0
        for right in range(n):
            currentSum += nums[right]
            while currentSum > goal:
                currentSum -= nums[left]
                left += 1
            counted += (right - left + 1)
        return counted