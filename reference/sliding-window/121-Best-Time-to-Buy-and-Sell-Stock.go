func maxProfit(prices []int) int {
    left := 0
    ans := 0
    for right, price := range prices {
        if prices[left] > price {
            left = right
        } else {
            ans = max(ans, price - prices[left])
        }
    }
    return ans
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}