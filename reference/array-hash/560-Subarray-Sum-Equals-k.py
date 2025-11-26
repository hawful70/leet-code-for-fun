class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        count = 0
        prefixSum = 0
        countMap = defaultdict(int)

        countMap[0] = 1

        for num in nums:
            prefixSum += num

            target = prefixSum - k
            if target in countMap:
                count += countMap[target]
            
            countMap[prefixSum] += 1
        return count