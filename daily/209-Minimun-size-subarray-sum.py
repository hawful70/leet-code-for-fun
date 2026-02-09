class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        n = len(nums)
        left = 0
        sum_num = 0
        res = float('inf')
        for right in range(n):
            sum_num += nums[right]

            while sum_num >= target:
                res = min(res, right - left + 1)
                sum_num -= nums[left]
                left += 1
        return res if res != float('inf') else 0