class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        # if len(nums) == 1:
        #     return
        # found = 0
        # for i in range(1, len(nums)):
        #     if nums[found] != nums[i] and nums[found] == 0:
        #         nums[found], nums[i] = nums[i], nums[found]
        #         found += 1
        #     if nums[found] != nums[i] and nums[i] == 0:
        #         found = i

        non_zero_pos = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[non_zero_pos] = nums[i]
                non_zero_pos += 1

        for i in range(non_zero_pos, len(nums)):
            nums[i] = 0

        
            
            
            
