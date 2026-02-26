from collections import defaultdict

class Solution:
    def subarraysWithKDistinct(self, nums: List[int], k: int) -> int:
        return self.atMost(nums, k) - self.atMost(nums, k - 1)
    def atMost(self, nums: List[int], k: int) -> int:
        count = defaultdict(int)
        left = res = 0
        n = len(nums)
        for right in range(n):
            count[nums[right]] += 1
            while len(count) > k:
                count[nums[left]] -= 1
                if count[nums[left]] == 0:
                    del count[nums[left]]
                left += 1
            res += right - left + 1
        return res