package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 3, 3, 4}
	fmt.Println(deleteAndEarn(nums))
}
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	sum := make([]int, maxVal+1)
	for _, num := range nums {
		sum[num] += num
	}
	return rob(sum)
}
func rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}
	return second
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
