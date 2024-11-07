class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0
        left, right = 0, 1
        while right < len(prices):
            if prices[left] < prices[right]:
                profit += prices[right] - prices[left]
            left += 1
            right += 1
        return profit
