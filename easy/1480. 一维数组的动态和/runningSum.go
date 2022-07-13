package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	fmt.Printf("%v", runningSum(nums))
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
