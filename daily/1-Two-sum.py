from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        map_sum = {}
        for index, num in enumerate(nums):
            remain = target - num
            if remain in map_sum:
                return [map_sum[remain], index]
            map_sum[num] = index
        return []
