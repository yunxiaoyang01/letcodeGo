package main

import "fmt"

func main() {
	nums := []int{1, 2, 5, 4, 7}
	fmt.Printf("%d\n", findLengthOfLCIS(nums))
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	i := 0
	j := 1
	ans := 0
	for j < len(nums) {
		if nums[j-1] >= nums[j] {
			i = j
		}
		j++
		ans = max(ans, j-i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
