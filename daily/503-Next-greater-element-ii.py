class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        n = len(nums)
        stacks = []
        ans = [-1] * n
        for i in range(2 * n):
            current_num = nums[i % n]

            while stacks and nums[stacks[-1]] < current_num:
                ans[stacks[-1]] = current_num
                stacks.pop()
            if i < n:
                stacks.append(i)
        
        return ans