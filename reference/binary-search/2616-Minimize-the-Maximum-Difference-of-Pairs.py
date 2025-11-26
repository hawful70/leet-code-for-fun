class Solution:
    def minimizeMax(self, nums: List[int], p: int) -> int:
        if len(nums) == 1:
            return 0
        nums.sort()
        # left, right = 0, 10**9
        left, right = 0, nums[-1] - nums[0]
        ans = 0
        while left <= right:
            mid = (left + right) // 2
            if self.isValidMaximumDifference(nums, p, mid):
                ans = mid
                right = mid - 1
            else:
                left = mid + 1
        return ans
    def isValidMaximumDifference(self, nums: List[int], p: int, threehold: int) -> bool:
        i, count = 0, 0
        while i < len(nums) - 1:
            if abs(nums[i] - nums[i + 1]) <= threehold:
                count += 1
                i += 2
            else:
                i += 1
            if count == p:
                return True
        return False
