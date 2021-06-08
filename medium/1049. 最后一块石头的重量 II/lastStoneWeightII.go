package main

func main() {

}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	half := sum/2
	dp := make([]int,half+1)
	dp[0] = 0
	for _, stone := range stones {
		for j:=half;j>=stone;j-- {
			dp[j] = max(dp[j],dp[j-stone]+stone)
		}
	}
	return sum-2*dp[half]
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}
