class Solution:
    def numberOfSubarrays(self, nums: List[int], k: int) -> int:
        return self.atMost(nums, k) - self.atMost(nums, k - 1)
    def atMost(self, nums: List[int], k: int) -> int:
        count, left, res = 0, 0, 0
        n = len(nums)
        for right in range(n):
            count += nums[right] % 2
            while count > k:
                count -= nums[left] % 2
                left += 1
            res += right - left + 1
        return res