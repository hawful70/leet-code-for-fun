class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        frequency_zero = defaultdict(int)
        count = 0
        left = 0
        n = len(nums)
        for right in range(n):
            if nums[right] != 1:
                frequency_zero[nums[right]] += 1
            while frequency_zero[nums[right]] > k:
                if nums[left] != 1:
                    frequency_zero[nums[left]] -= 1
                left += 1
            count = max(count, right - left + 1)
        return count