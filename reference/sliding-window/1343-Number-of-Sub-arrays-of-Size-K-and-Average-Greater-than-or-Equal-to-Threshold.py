class Solution:
    def numOfSubarrays(self, arr: List[int], k: int, threshold: int) -> int:
        ans = 0
        cur_sum = sum(arr[:k-1])
        for l in range(len(arr) - k + 1):
            cur_sum += arr[l + k - 1]
            if (cur_sum / k) >= threshold:
                ans += 1
            cur_sum -= arr[l]
        return ans