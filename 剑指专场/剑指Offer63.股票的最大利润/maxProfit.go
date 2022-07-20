package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := 0
	min := prices[0]
	for _, price := range prices {
		if min > price {
			min = price
		}
		if price-min > max {
			max = price - min
		}
	}
	return max
}
