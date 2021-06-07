package main

import "fmt"

func main() {
	nums := []int{23, 2, 6, 4, 7}
	k := 6

	fmt.Println(checkSubarraySum(nums, k))
}

func checkSubarraySum(nums []int, k int) bool {
	m := len(nums)
	if m < 2 {
		return false
	}
	result := map[int]int{
		0: -1,
	}
	re := 0
	for i, num := range nums {
		re = (re + num) % k

		if pre, ok := result[re]; ok {
			if i-pre >= 2 {
				return true
			}
		} else {
			result[re] = i
		}
	}
	return false
}
