class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        left = 0
        n = len(prices)
        res = 0
        for right in range(1, n):
            current_profit = prices[right] - prices[left]
            if current_profit >= 0:
                res = max(res, current_profit)
            else:
                left = right
        return res