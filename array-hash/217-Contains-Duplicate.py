class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        duplicate_map = {}
        for i in range(len(nums)):
            if nums[i] in duplicate_map:
                return True
            duplicate_map[nums[i]] = 1
        return False