package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	stack := []int{}
	for i := 0; i < 2*n-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Printf("%d", nextGreaterElements(nums))
}
