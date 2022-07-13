package main

import "fmt"

func main() {
	prices := []int{1, 7, 5, 3, 6, 4}
	fmt.Printf("prices: %d", maxProfit(prices))
}
func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	mid := 0
	max := 0
	min := prices[0]
	for _, v := range prices {
		if v < min {
			min = v
		}
		mid = v - min
		if mid > max {
			max = mid
		}
	}
	return max
}
