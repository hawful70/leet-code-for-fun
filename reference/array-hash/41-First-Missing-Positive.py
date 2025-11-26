class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        n = len(nums)
        for i in range(n):
            index = nums[i] - 1
            while 0 < nums[i] <= n and nums[index] != nums[i]:
                nums[index], nums[i] = nums[i], nums[index]
                index = nums[i] - 1

        for i in range(n):
            if nums[i] != i + 1:
                return i + 1
        return n + 1
