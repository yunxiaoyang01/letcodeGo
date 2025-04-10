package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
