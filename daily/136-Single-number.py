from collections import defaultdict
from typing import List

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        map_num = defaultdict(int)
        for num in nums:
           map_num[num] += 1

        for num in map_num:
            if map_num[num] != 2:
                return num
        return 0