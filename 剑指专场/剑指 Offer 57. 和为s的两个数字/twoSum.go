package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Printf("%v", twoSum(nums, 9))
}

// 题解 题目写了递增数组，设定左指针为0下标，右指针为len(nums)-1下标
// 如果 左指针+右指针的值大于target 证明右边大了，右边减一
// 反之 左面小了，左面加+1
// 最终 找到target
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]+nums[right] < target {
			left = left + 1
		} else if nums[left]+nums[right] > target {
			right = right - 1
		} else {
			break
		}
	}
	return []int{nums[left], nums[right]}
}
