class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        if len(nums) == 1 and nums[0] == target:
            return [0, 0]
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                if nums[left] != nums[mid]:
                    left += 1
                elif nums[right] != nums[mid]:
                    right -= 1
                else:
                    return [left, right]
            elif nums[mid] > target:
                right = mid - 1
            else:
                left = mid + 1
        return [-1, -1]
