package easy

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func containsDuplicate1(nums []int) bool {
	data := map[int]bool{}
	for _, x := range nums {
		if _, ok := data[x]; ok {
			return true
		}
		data[x] = true
	}
	return false
}
