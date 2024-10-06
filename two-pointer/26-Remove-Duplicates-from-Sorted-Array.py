class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        found = 0
        for i in range(1, len(nums)):
            if nums[i] != nums[found]:
                found += 1
                nums[found], nums[i] = nums[i], nums[found]
        return found + 1