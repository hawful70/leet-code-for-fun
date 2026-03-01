class Solution:
    def shortestSubarray(self, nums: List[int], k: int) -> int:
        n = len(nums)
        res = float('inf')
        # Prefix sum array: P[i] is sum of first i elements
        P = [0] * (n + 1)
        for i in range(n):
            P[i+1] = P[i] + nums[i]
            
        dq = deque() # Will store indices in increasing order of P[i]
        
        for i in range(n + 1):
            # 1. OPTIMIZE: If we found a valid subarray, try to shrink from the left
            while dq and P[i] - P[dq[0]] >= k:
                res = min(res, i - dq.popleft())
                
            # 2. MAINTAIN MONOTONICITY: If current P[i] is smaller than the last one,
            # the last one is useless because P[i] is a "better" starting point (smaller P[i] makes P[j]-P[i] larger)
            while dq and P[i] <= P[dq[-1]]:
                dq.pop()
                
            dq.append(i)
            
        return res if res != float('inf') else -1