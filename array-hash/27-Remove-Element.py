class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        left, right = 0, len(nums) - 1
    
        while left < right:
            if nums[left] == val:
                if nums[left] != nums[right]:
                    nums[left], nums[right] = nums[right], nums[left]
                    left += 1
                right -= 1
            else:
                left += 1

        k = 0
        for i in range(len(nums)):
            if nums[i] != val:
                k += 1
            else:
                break
        
        return k