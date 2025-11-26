class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0
        left = 0
        for right in range(1, len(prices)):
            if prices[left] > prices[right]:
                left = right
            else:
                currentProfit = prices[right] - prices[left]
                profit = max(profit, currentProfit)
        return profit