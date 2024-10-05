class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        \\\
        Do not return anything, modify nums in-place instead.
        \\\
        if len(nums) == 1:
            return
        found = 0
        for i in range(1, len(nums)):
            if nums[found] != nums[i] and nums[found] == 0:
                nums[found], nums[i] = nums[i], nums[found]
                found += 1
            if nums[found] != nums[i] and nums[i] == 0:
                found = i
        
            
            
            
