package easy

/**
在这种情况下，我们可以简单地继续在斜坡上爬升并持续增加从连续交易中获得的利润，
而不是在谷之后寻找每个峰值。最后，我们将有效地使用峰值和谷值，
但我们不需要跟踪峰值和谷值对应的成本以及最大利润，但我们可以直接继续增加加数组的连续数字之间的差值，
如果第二个数字大于第一个数字，我们获得的总和将是最大利润。这种方法将简化解决方案。

不明觉厉，先记下有时间研究
*/
func maxProfitTwo(prices []int) int {
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
