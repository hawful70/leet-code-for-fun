class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) <= 2:
            return len(nums)

        found = 2
        for i in range(2, len(nums)):
            if nums[i] != nums[found - 2]:
                nums[i], nums[found] = nums[found], nums[i]
                found += 1
        return found