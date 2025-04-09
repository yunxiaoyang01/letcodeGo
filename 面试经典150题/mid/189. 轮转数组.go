package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Printf("%v\n", nums[len(nums)-3:])
	rotate(nums, k)
	fmt.Printf("%v", nums)
}

func rotate(nums []int, k int) {
	n := len(nums)
	l := n - k%len(nums)
	right := nums[l:]
	left := nums[:l]
	right = append(right, left...)
	copy(nums, right)
}
