class Solution:
    def arrangeCoins(self, n: int) -> int:
        left, right = 1, n
        ans = 0

        # math: (n / 2) * (n + 1) == 1 + 2 + 3 + ... + n
        while left <= right:
            mid = (left + right) // 2
            coins = (mid / 2) * (mid + 1)
            if coins <= n:
                ans = max(ans, mid)
                left = mid + 1
            else:
                right = mid - 1
        return ans