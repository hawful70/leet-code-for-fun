class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = defaultdict(int)
        frequency = [[] for _ in range(len(nums) + 1)]
        
        for n in nums:
            count[n] += 1
        ans = []
        for key, value in count.items():
            frequency[value].append(key)
        for i in range(len(frequency) - 1, -1, -1):
            for n in frequency[i]:
                ans.append(n)
                if len(ans) == k:
                    return ans
        return ans