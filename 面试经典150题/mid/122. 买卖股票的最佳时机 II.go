package main

func maxProfit(prices []int) int {
	total := 0
	for i := 0; i < len(prices)-1; i++ {
		total += max(prices[i+1]-prices[i], 0)

	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
