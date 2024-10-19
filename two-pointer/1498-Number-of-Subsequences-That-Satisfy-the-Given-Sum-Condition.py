class Solution(object):
    def numSubseq(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        nums.sort()
        MOD = 10**9 + 7
        ans = 0
        left, right = 0, len(nums) - 1
        while left <= right:
            if nums[left] + nums[right] <= target:
                ans += pow(2, right - left)
                ans %= MOD
                left += 1
            else:
                right -= 1
        return ans
