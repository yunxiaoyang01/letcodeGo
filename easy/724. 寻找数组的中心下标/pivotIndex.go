package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Printf("%v", pivotIndex(nums))
}

// 解法 当遍历到第i个元素的时候，i左侧的所有元素之和等于i 右侧的元素之和 所以 sumL+numi+ sumR =total
// 即 2*sum+numi == total

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}
