package main

// 解法一 完全暴力
func majorityElement(nums []int) int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

// 解法二 消杀法
func majorityElementV2(nums []int) int {
	var winners = nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if winners == nums[i] {
			count++
		} else if count == 0 {
			winners = nums[i]
			count++
		} else {
			count--
		}
	}
	return winners
}
