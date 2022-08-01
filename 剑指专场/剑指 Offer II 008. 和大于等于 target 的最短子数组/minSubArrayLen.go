package main

import "math"

func main() {

}

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	minNum := math.MaxInt32
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			minNum = min(minNum, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if minNum == math.MaxInt32 {
		return 0
	}
	return minNum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
