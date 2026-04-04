class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if len(nums) > 1:
            left, right = 0, len(nums) - 1
            k = k % len(nums)
            self.reverseArray(left, right, nums)
            self.reverseArray(0, k - 1, nums)
            self.reverseArray(k, right, nums)

    def reverseArray(self, left: int, right: int, nums: List[int]) -> None:

        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1