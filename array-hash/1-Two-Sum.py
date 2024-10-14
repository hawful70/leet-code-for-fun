class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hash_sum = {}
        for i in range(len(nums)):
            result = target - nums[i]
            if result in hash_sum:
                return [i, hash_sum[result]]
            hash_sum[nums[i]] = i
        return []