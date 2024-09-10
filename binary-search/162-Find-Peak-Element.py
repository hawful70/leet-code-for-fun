class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 0

        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if (mid == right and nums[mid] > nums[mid - 1]) or (nums[mid] > nums[mid + 1] and nums[mid] > nums[mid - 1]):
                return mid
            elif nums[mid] < nums[mid + 1]:
                left = mid + 1
            else:
                right = mid - 1
        return mid