class Solution:
    def sortColors(self, nums: List[int]) -> None:
        current, left, right = 0, 0, len(nums) - 1
        while current <= right:
            if nums[current] == 0:
                nums[current], nums[left] = nums[left], nums[current]
                left = left + 1
                current = current + 1
            elif nums[current] == 2:
                nums[current], nums[right] = nums[right], nums[current]
                right = right - 1
            else:
                current = current + 1
        