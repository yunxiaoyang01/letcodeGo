package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] == nums[l] && nums[mid] == nums[r] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			if nums[l] < target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target < nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 2
	fmt.Println(search(nums, target))
}
