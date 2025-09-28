from bisect import bisect_right
from typing import List


class Solution:
    def maxProfitAssignment(self, difficulty: List[int], profit: List[int], worker: List[int]) -> int:
        jobs = sorted(zip(difficulty, profit))
        best_profit = []

        for diff, prof in jobs:
            if not best_profit:
                best_profit.append(prof)
            else:
                best_profit.append(max(best_profit[-1], prof))

        difficulties = [diff for diff, _ in jobs]
        total = 0

        for capacity in worker:
            idx = bisect_right(difficulties, capacity) - 1
            if idx >= 0:
                total += best_profit[idx]

        return total
