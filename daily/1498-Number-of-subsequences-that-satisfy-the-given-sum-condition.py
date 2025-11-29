from typing import List


class Solution:
    def numSubseq(self, nums: List[int], target: int) -> int:
        if not nums:
            return 0

        mod = 10**9 + 7
        nums.sort()
        n = len(nums)

        pow2 = [1] * n
        for i in range(1, n):
            pow2[i] = (pow2[i - 1] * 2) % mod

        ans = 0
        for left in range(n):
            limit = target - nums[left]
            if limit < nums[left]:
                break

            lo, hi = left, n - 1
            while lo <= hi:
                mid = (lo + hi) // 2
                if nums[mid] <= limit:
                    lo = mid + 1
                else:
                    hi = mid - 1

            right = hi
            if right >= left:
                ans = (ans + pow2[right - left]) % mod

        return ans
