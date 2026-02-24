from collections import deque

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        res = []
        n = len(nums)
        max_deque = deque()
        # for i in range(k):
        #     while max_deque and nums[max_deque[-1]] < nums[i]:
        #         max_deque.pop()
        #     max_deque.append(i)
        # res.append(nums[max_deque[0]])

        for i in range(n):
            if max_deque and i - max_deque[0] >= k:
                max_deque.popleft()

            while max_deque and nums[max_deque[-1]] < nums[i]:
                   max_deque.pop()
            max_deque.append(i)
            
            if i >= k - 1:
                res.append(nums[max_deque[0]])

        # return res if k != 1 else nums
        return res