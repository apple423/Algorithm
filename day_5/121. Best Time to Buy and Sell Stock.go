package day_5

func MaxProfit(prices []int) int {
	profit := 0
	min := prices[0]

	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}

		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}
