func maxProfit(prices []int) int {
    profit := 0
    left := 0
    for right, price := range prices {
        if prices[left] > price {
            left = right
        } else {
            currentProfit := price - prices[left]
            profit = max(profit, currentProfit)
        }
    }
    return profit
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}