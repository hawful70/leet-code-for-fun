class Solution:
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        if len(nums) == 1:
            return nums
        
        low, high = 0, 1
        while low <= high < len(nums):
            if nums[low] % 2 == 0:
                low += 1
            if nums[low] % 2 != 0 and nums[high] % 2 == 0:
                nums[low], nums[high] = nums[high], nums[low]
                low += 1
            high += 1
        return nums