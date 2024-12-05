class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        min_len = len(nums)

        current_sum = 0
        left, right = 0, 0
        isFound = False

        while right < len(nums):
            current_sum += nums[right]

            while current_sum >= target:
                min_len = min(min_len, right - left + 1)
                current_sum -= nums[left]
                isFound = True
                left += 1
            right += 1

        return min_len if isFound else 0