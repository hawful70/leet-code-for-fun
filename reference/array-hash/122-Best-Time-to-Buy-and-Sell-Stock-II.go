func maxProfitV2(prices []int) int {
    profit := 0
    left := 0
    sum_profit := 0
    dp := make([]int, len(prices))
    for right := 1; right < len(prices); right++ {
        if prices[right] > prices[left] && prices[right] >= prices[right - 1] {
            profit = prices[right] - prices[left]
            dp[left] = profit
        } else {
            left = right
        }
    }
    for _, price := range dp {
        sum_profit += price
    }
    return sum_profit
}

func maxProfit(prices []int) int {
    profit := 0
    left, right := 0, 1
    for right < len(prices) {
        if prices[left] < prices[right] {
            profit += prices[right] - prices[left]
        }
        left++
        right++
    }
    return profit
}