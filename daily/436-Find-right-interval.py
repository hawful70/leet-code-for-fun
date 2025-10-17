from bisect import bisect_left
from typing import List


class Solution:
    def findRightInterval(self, intervals: List[List[int]]) -> List[int]:
        n = len(intervals)
        starts = sorted((interval[0], idx) for idx, interval in enumerate(intervals))
        start_values = [start for start, _ in starts]

        result = [-1] * n
        for i, (_, end) in enumerate(intervals):
            pos = bisect_left(start_values, end)
            if pos < n:
                result[i] = starts[pos][1]

        return result
