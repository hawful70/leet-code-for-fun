class Solution:
    def splitArray(self, nums: List[int], k: int) -> int:
        # Search space:
        # - Minimum possible largest sum = max element
        # - Maximum possible largest sum = sum of all elements
        left, right = max(nums), sum(nums)

        while left < right:
            mid = (left + right) // 2

            # Check feasibility: can we split into <= k subarrays
            # such that each subarray sum <= mid?
            if self.canSplit(nums, mid, k):
                right = mid      # try smaller maximum sum
            else:
                left = mid + 1   # need larger maximum sum

        return left

    def canSplit(self, nums: List[int], limit: int, k: int) -> bool:
        count = 1
        current_sum = 0

        for num in nums:
            current_sum += num
            if current_sum > limit:
                count += 1
                current_sum = num
                if count > k:
                    return False

        return True
