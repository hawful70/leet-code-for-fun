from typing import List


class Solution:
    def minCostRearrangingFruits(self, basket1: List[int], basket2: List[int]) -> int:
        freq1 = {}
        freq2 = {}
        all_freq = {}
        for x in basket1:
            freq1[x] = freq1.get(x, 0) + 1
            all_freq[x] = all_freq.get(x, 0) + 1
        for x in basket2:
            freq2[x] = freq2.get(x, 0) + 1
            all_freq[x] = all_freq.get(x, 0) + 1

        for v in all_freq.values():
            if v % 2 != 0:
                return -1

        swap_from1 = []
        swap_from2 = []
        for fruit in all_freq.keys():
            diff = freq1.get(fruit, 0) - freq2.get(fruit, 0)
            if diff > 0:
                swap_from1.extend([fruit] * (diff // 2))
            elif diff < 0:
                swap_from2.extend([fruit] * ((-diff) // 2))

        swap_from1.sort()
        swap_from2.sort(reverse=True)

        min_fruit = min(min(basket1), min(basket2))
        cost = 0
        for a, b in zip(swap_from1, swap_from2):
            direct = min(a, b)
            indirect = 2 * min_fruit
            cost += min(direct, indirect)
        return cost

