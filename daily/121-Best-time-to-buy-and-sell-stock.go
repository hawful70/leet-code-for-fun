package daily

func maxProfit(prices []int) int {
	left := 0
	n := len(prices)
	res := 0
	for right := 1; right < n; right++ {
		currentProfit := prices[right] - prices[left]
		if currentProfit >= 0 {
			res = max(res, currentProfit)
		} else {
			left = right
		}
	}
	return res
}
