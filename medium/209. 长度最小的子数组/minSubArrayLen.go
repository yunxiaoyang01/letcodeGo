package main

import "math"

// func minSubArrayLen(s int, nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	res := math.MaxInt32
// 	left, right := 0, 0
// 	sum := 0
// 	for right < n {
// 		sum += nums[right]
// 		for sum >= s {
// 			res = min(res, right-left+1)
// 			sum -= nums[left]
// 			left++
// 		}
// 		right++
// 	}
// 	if res == math.MaxInt32 {
// 		return 0
// 	}
// 	return res
// }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	left, right := 0, 0
	sum := 0
	ans := math.MaxInt32
	for right < l {
		sum += nums[right]
		for sum >= s {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
