package main

import "sort"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	println(searchInsert(nums, target))
}

/**
35. 搜索插入位置
*/
func searchInsert(nums []int, target int) int {
	k := -1
	nums = append(nums, target)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			k = i
			break
		}
	}
	return k
}
