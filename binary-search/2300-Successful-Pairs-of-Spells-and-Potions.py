class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        n = len(potions)
        potions.sort()
        ans = []

        for i in range(len(spells)):
            lower_bound_index = self.binarySearch(spells[i], potions, success)
            lower_bound_index = n - lower_bound_index if lower_bound_index != -1 else 0
            ans.append(lower_bound_index)
        return ans

    def binarySearch(self, spell: int, potions: List[int], target: int) -> int:
        left, right = 0, len(potions) - 1
        result = -1
        while left <= right:
            mid = (left + right) // 2
            if spell * potions[mid] >= target:
                result = mid
                right = mid - 1
            else:
                left = mid + 1
        return result