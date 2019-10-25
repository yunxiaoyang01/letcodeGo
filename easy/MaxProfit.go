package easy

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	print(maxProfit(prices))
}

func maxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			temp := prices[j] - prices[i]
			if temp > sum {
				sum = temp
			}
		}

	}
	return sum
}

//动态规划优化原答案
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	thisProfit := 0
	maxProfit := 0
	min := prices[0]
	for _, v := range prices {
		if v < min {
			min = v
		}
		thisProfit = v - min
		if thisProfit > maxProfit {
			maxProfit = thisProfit
		}
	}
	return maxProfit
}
