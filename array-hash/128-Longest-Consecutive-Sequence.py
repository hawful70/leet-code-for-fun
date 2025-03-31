class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        numSet = set(nums)
        maxLength = 0

        for num in numSet:
            if num - 1 not in numSet:
                current = 0
                while num + current in numSet:
                    current += 1
                maxLength = max(maxLength, current)
        return maxLength